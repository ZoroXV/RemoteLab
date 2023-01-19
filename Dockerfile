FROM golang:1.18.3-bullseye as builder
WORKDIR /app
RUN apt-get update && apt-get install -y libusb-1.0-0-dev
COPY ./src /app
RUN ["go", "get", "github.com/google/gousb"]
RUN ["go", "get", "github.com/google/gousb/usbid"]
RUN ["go", "get", "github.com/citilinkru/libudev"]
RUN ["go", "get", "golang.org/x/exp/slices"]
RUN ["go", "build"]

FROM debian:bullseye as release_server
RUN apt-get update \
    && apt-get install -y \
    stlink-tools \
    curl
RUN curl -fsSL https://raw.githubusercontent.com/arduino/arduino-cli/master/install.sh | BINDIR=/usr/local/bin sh && arduino-cli core install arduino:avr
WORKDIR /app
COPY --from=builder /app/remotelab ./
COPY --from=builder /app/cli/remotelab.py ./
CMD ["./remotelab"]

FROM python:latest as release_gui
WORKDIR /app
COPY ./src/frontend /app
RUN ["pip3", "install", "-r", "requirements.txt"]
CMD ["python3", "app.py"]
FROM golang:1.18.3-bullseye as builder
WORKDIR /app
COPY ./src /app
RUN ["go", "build"]

FROM debian:bullseye as release
RUN apt-get update \
    && apt-get install -y \
    stlink-tools \
    curl
RUN curl -fsSL https://raw.githubusercontent.com/arduino/arduino-cli/master/install.sh | BINDIR=/usr/local/bin sh && arduino-cli core install arduino:avr
WORKDIR /app
COPY --from=builder /app/remotelab ./
CMD ["./remotelab"]

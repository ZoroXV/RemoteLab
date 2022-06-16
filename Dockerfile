FROM golang:1.18.3-bullseye AS builder
WORKDIR /app
COPY ./src /app
RUN go build

FROM debian:bullseye AS runner
RUN apt-get update \
    && apt-get install -y \
    stlink-tools \
    curl
RUN curl -fsSL https://raw.githubusercontent.com/arduino/arduino-cli/master/install.sh | BINDIR=/usr/local/bin sh
WORKDIR /root
COPY --from=builder /app/remotelab ./
CMD ["./remotelab"]

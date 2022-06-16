FROM debian:bullseye
RUN apt-get update \
    && apt-get install -y \
    stlink-tools \
    curl
RUN curl -fsSL https://raw.githubusercontent.com/arduino/arduino-cli/master/install.sh | BINDIR=/usr/local/bin sh
WORKDIR /
COPY remotelab ./
CMD ["./remotelab"]

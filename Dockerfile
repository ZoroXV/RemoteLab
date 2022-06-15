FROM ubuntu:22.04

LABEL maintainer="RemoteLab"
LABEL version="0.1"

RUN apt-get update \
    && apt-get install -y\
        build-essential \
        stlink-tools \
        curl \
        wget
RUN rm -rf /usr/local/go && wget https://go.dev/dl/go1.18.2.linux-amd64.tar.gz && tar -C /usr/local -xzf go1.18.2.linux-amd64.tar.gz
RUN curl -fsSL https://raw.githubusercontent.com/arduino/arduino-cli/master/install.sh | sh

COPY ./src /app
WORKDIR /app

RUN /usr/local/go/bin/go build
CMD ./remotelab

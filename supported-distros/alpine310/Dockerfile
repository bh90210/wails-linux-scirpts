FROM alpine:3.10
LABEL wails="removeall" maintainer="ktc@protonmail.com"

RUN apk update && \
apk add gtk+3.0 wget git webkit2gtk-dev npm build-base && \
 wget https://dl.google.com/go/go1.12.7.linux-amd64.tar.gz && tar -C /usr/local -xzf go1.12.7.linux-amd64.tar.gz && \
  && \
apk -v cache clean

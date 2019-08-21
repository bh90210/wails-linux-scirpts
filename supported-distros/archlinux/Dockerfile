FROM archlinux/base
LABEL wails="removeall" maintainer="ktc@protonmail.com"

RUN pacman -Syu --noconfirm git curl wget nodejs npm gtk3 webkit2gtk gcc pkgconf && \
 wget https://dl.google.com/go/go1.12.7.linux-amd64.tar.gz && tar -C /usr/local -xzf go1.12.7.linux-amd64.tar.gz

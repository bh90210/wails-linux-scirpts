FROM voidlinux/voidlinux-musl:latest
LABEL wails="removeall" maintainer="ktc@protonmail.com"

RUN xbps-install -S bash
ENTRYPOINT [ "/bin/bash" ]
#RUN xbps-install -Sy git curl wget base-devel pkg-config nodejs gtk+3-devel webkit2gtk-devel && \
#wget https://dl.google.com/go/go1.12.7.linux-amd64.tar.gz && tar -C /usr/local -xzf go1.12.7.linux-amd64.tar.gz

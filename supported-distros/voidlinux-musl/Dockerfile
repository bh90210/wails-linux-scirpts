FROM voidlinux/voidlinux-musl:latest
LABEL wails="removeall" maintainer="ktc@protonmail.com"

RUN xbps-install -Syu
RUN xbps-install -yu
RUN xbps-install -y go git curl wget base-devel pkg-config nodejs gtk+3-devel webkit2gtk-devel
ENTRYPOINT ["/bin/sh"]

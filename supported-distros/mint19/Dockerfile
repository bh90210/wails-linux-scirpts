FROM linuxmintd/mint19-amd64
LABEL wails="removeall" maintainer="ktc@protonmail.com"

RUN apt-get update && \
 apt-get install -y git curl wget libgtk-3-dev libwebkit2gtk-4.0-dev build-essential npm nodejs && \
 wget https://dl.google.com/go/go1.12.7.linux-amd64.tar.gz && tar -C /usr/local -xzf go1.12.7.linux-amd64.tar.gz && \
 rm -rf /var/lib/apt/lists/*

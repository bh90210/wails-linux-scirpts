ARG DISTRO
FROM wail-$DISTRO
ENV PATH=$PATH:/usr/local/go/bin:/root/go/bin
VOLUME /wails
ADD wails /wails
ADD ./wails /
ENTRYPOINT ["/bin/bash", "-c", "echo Hello, $name"]

ARG DISTRO
FROM wails-$DISTRO
ENV PATH "$PATH:/usr/local/go/bin"
ADD wails /usr/local/go/bin
COPY source /wails
ENTRYPOINT ["/bin/bash"]

FROM  opensuse/tumbleweed
LABEL wails="removeall" maintainer="ktc@protonmail.com"

RUN zypper addrepo https://download.opensuse.org/repositories/network:utilities/openSUSE_Factory/network:utilities.repo | echo 'a' && \
zypper refresh && \
zypper --non-interactive install --type pattern devel_basis && \
zypper --non-interactive install git-core wget libgtk-3-0 libwebkit2gtk-4_0-37 nodejs10 npm && \
wget https://dl.google.com/go/go1.12.7.linux-amd64.tar.gz && tar -C /usr/local -xzf go1.12.7.linux-amd64.tar.gz && \
zypper clean

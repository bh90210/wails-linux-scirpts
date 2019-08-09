FROM centos:7
LABEL wails="removeall" maintainer="ktc@protonmail.com"

RUN yum -y update && \
 yum -y install curl wget  gcc gcc-c++ make pkgconf-pkg-config nodejs npm gtk3-devel webkitgtk3-devel && \
 wget https://dl.google.com/go/go1.12.7.linux-amd64.tar.gz && tar -C /usr/local -xzf go1.12.7.linux-amd64.tar.gz

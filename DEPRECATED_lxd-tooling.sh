#!/bin/bash

# first 
# sudo chmod +x script-name.sh
#
# how to use

# create a container (optionally pass all to create every supported distro)
# ./lxd-tooling.sh craete centos7 (all)
#                         lxd os

# ./lxd-tooling.sh delete centos7 (all)
#                         lxd os

COMMAND=$1
DISTRO=$2
GOVERSION='go1.12.7.linux-amd64.tar.gz'

debian9(){
local CONTAINERNAME='wails-debian9-test'
if [ $1 = "create" ]
then
    sudo lxc launch images:debian/9 ${CONTAINERNAME}
    sleep 20
    sudo lxc exec ${CONTAINERNAME} --  sh -c "apt-get update && apt-get upgrade && apt-get install git curl wget libgtk-3-dev libwebkit2gtk-4.0-dev build-essential && curl -sL https://deb.nodesource.com/setup_12.x | bash - && apt-get install -y nodejs && wget https://dl.google.com/go/${GOVERSION} && tar -C /usr/local -xzf ${GOVERSION}"
    sudo lxc stop ${CONTAINERNAME}
elif [ $1 = "delete" ]
then
    sudo lxc delete ${CONTAINERNAME} --force
fi
}

centos7(){
local CONTAINERNAME='wails-centos7-test'
if [ $1 = "create" ]
then
    sudo lxc launch images:centos/7 ${CONTAINERNAME}
    sleep 20
    sudo lxc exec ${CONTAINERNAME} --  sh -c "yum update && yum install curl wget gcc gcc-c++ make pkgconf-pkg-config  gtk3-devel webkitgtk3-devel && curl -sL https://deb.nodesource.com/setup_12.x | bash - && yum install -y nodejs && wget https://dl.google.com/go/${GOVERSION} && tar -C /usr/local -xzf ${GOVERSION}"
    sudo lxc stop ${CONTAINERNAME}
elif [ $1 = "delete" ]
then
    sudo lxc delete ${CONTAINERNAME} --force
fi
}

fedora30(){
local CONTAINERNAME='wails-fedora30-test'
if [ $1 = "create" ]
then
    sudo lxc launch images:fedora/30 ${CONTAINERNAME}
    sleep 20
    sudo lxc exec ${CONTAINERNAME} --  sh -c "apt-get update && apt-get upgrade && apt-get install git curl wget libgtk-3-dev libwebkit2gtk-4.0-dev build-essential && curl -sL https://deb.nodesource.com/setup_12.x | bash - && apt-get install -y nodejs && wget https://dl.google.com/go/${GOVERSION} && tar -C /usr/local -xzf ${GOVERSION}"
    sudo lxc stop ${CONTAINERNAME}
elif [ $1 = "delete" ]
then
    sudo lxc delete ${CONTAINERNAME} --force
fi
}

archlinux(){
local CONTAINERNAME='wails-archlinux-test'
if [ $1 = "create" ]
then
    sudo lxc launch images:archlinux ${CONTAINERNAME}
    sleep 20
    sudo lxc exec ${CONTAINERNAME} --  sh -c "apt-get update && apt-get upgrade && apt-get install git curl wget libgtk-3-dev libwebkit2gtk-4.0-dev build-essential && curl -sL https://deb.nodesource.com/setup_12.x | bash - && apt-get install -y nodejs && wget https://dl.google.com/go/${GOVERSION} && tar -C /usr/local -xzf ${GOVERSION}"
    sudo lxc stop ${CONTAINERNAME}
elif [ $1 = "delete" ]
then
    sudo lxc delete ${CONTAINERNAME} --force
fi
}


all(){
debian9 $1
fedora30 $1
centos7 $1
archlinux $1
}

if [ ${COMMAND} = "create" ]
then
    ${DISTRO} create
elif [ ${COMMAND} = "delete" ]
then
    ${DISTRO} delete
fi


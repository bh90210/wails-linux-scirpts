#!/bin/bash

# how to use
# ./ansible-wails-branch-test-lxc.sh https://github.com/bh90210/wails.git develop centos7
#                                    git                                  branch  lxd os

RED='\033[0;31m'
NC='\033[0m' # No Color
GIT=$1
BRANCH=$2

# functions
# fedora30
fedora30(){
    echo -e "starting fedora30 container"
    sudo lxc start wails-fedora-test
    sleep 2

    runFunction "wails-fedora-test"

    echo -e "stoping fedora30 container"
    sudo lxc stop wails-fedora-test
    sleep 2
}

# debian9
debian9() {
    echo -e "starting debian9 container"
    sudo lxc start wails-debian-test
    sleep 2

    runFunction "wails-debian-test"

    echo -e "stoping debian9 container"
    sudo lxc stop wails-debian-test
    sleep 2
}

# archlinux
archlinux(){
    echo -e "starting archlinux container"
    sudo lxc start wails-archlinux-test
    sleep 2

    runFunction "wails-archlinux-test"

    echo -e "stoping archlinux container"
    sudo lxc stop wails-archlinux-test
    sleep 2
}


# centos7
centos7(){
    echo -e "starting centos7 container"
    sudo lxc start wails-centos-test
    sleep 2

    runFunction "wails-centos-test"

    echo -e "stoping centos7 container"
    sudo lxc stop wails-centos-test
    sleep 2
}

runFunction(){
    local DISTRO=$1
    echo -e "replacing previous wails installation with specified git/branch"
    sudo lxc exec ${DISTRO} -- rm /root/go/bin/wails
    sudo lxc exec ${DISTRO} -- rm -r /root/wails
    sudo lxc exec ${DISTRO} -- sh -c "git clone -b ${BRANCH} ${GIT} /root/wails && cd wails/cmd/wails && export PATH=$PATH:/usr/local/go/bin && go install"
    echo -e "wails branch ${BRANCH} from  ${GIT} was go installed succesfully"
    echo -e "init & build wails project? (Y) (anything else returns active container bash)"
    read input
    if [[ $input == "Y" || $input == "y" ]]; then
        sudo lxc exec ${DISTRO} -- rm -r /root/project
        sudo lxc exec ${DISTRO} -- sh -c "cd ~ && export PATH=$PATH:/usr/local/go/bin:/root/go/bin && wails init"
        sudo lxc exec ${DISTRO} -- sh -c "export PATH=$PATH:/usr/local/go/bin:/root/go/bin  && cd /root/project && wails build -d"
        sudo lxc exec ${DISTRO} -- /bin/bash
    else
        sudo lxc exec ${DISTRO} -- /bin/bash
    fi
}

if [ $3 = "all" ]
then
    debian9
    centos7
    fedora30
    archlinux
elif [ $3 = "debian9" ]
then
    debian9
elif [ $3 = "centos7" ]
then
    centos7
elif [ $3 = "fedora30" ]
then
    fedora30
elif [ $3 = "archlinux" ]
then
    archlinux
fi

#!/bin/bash

# how to use
# ./ansible-wails-branch-test-lxc.sh https://github.com/bh90210/wails.git develop centos7
#                                    git                                  branch  lxd os

RED='\033[0;31m'
NC='\033[0m' # No Color
GIT=$1
BRANCH=$2

sudo lxc start wails-centos-test
sudo lxc start wails-fedora-test
sudo lxc start wails-archlinux-test
sudo lxc start wails-debian-test
sleep 4

# functions
# fedora30
fedora30(){
    runFunction "wails-fedora-test"
}

# debian9
debian9() {
    runFunction "wails-debian-test"
}

# archlinux
archlinux(){
    runFunction "wails-archlinux-test"
}


# centos7
centos7(){
    runFunction "wails-centos-test"
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
# after testing part
closing(){
    echo -e "stoping all containers"
    sudo lxc stop wails-debian-test
    sudo lxc stop wails-centos-test
    sudo lxc stop wails-fedora-test
    sudo lxc stop wails-archlinux-test
}

if [ $3 = "all" ]
then
    debian9
    centos7
    fedora30
    archlinux
    closing
    
elif [ $3 = "debian9" ]
then
    debian9
    closing
    
elif [ $3 = "centos7" ]
then
    centos7
    closing
    
elif [ $3 = "fedora30" ]
then
    fedora30
    closing
    
elif [ $3 = "archlinux" ]
then
    archlinux
    closing
fi

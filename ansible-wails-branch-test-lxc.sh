#!/bin/bash

# functions
# fedora30
fedora30(){
    echo -e "starting fedora 30 lxc"
    sudo lxc start wails-fedora-test
    echo -e "exec in the bin/bash of the container"
    sudo lxc exec wails-fedora-test  -- /bin/bash
    
    initTestingWailsSetup
    
    echo -e "is wails init & wails build -d enough? (y/n)"
    read input
    if [ $input = "y" ]
    then
        buildWails
    elif [ $input = "n" ]
    then
        enterCommand
    fi
}

# centos7
centos7(){
    echo -e "starting centos 7 lxc"
    sudo lxc start wails-centos-test
    
    initTestingWailsSetup $1 $2

    echo -e "is wails init & wails build -d enough? (y/n)"
    read input
    if [ $input = "y" ]
    then
        buildWails
    elif [ $input = "n" ]
    then
        enterCommand
    fi
}

# debian9
debian9() {
    echo -e "starting debian 9 lxc"
    sudo lxc start wails-debian-test
    echo -e "exec in the bin/bash of the container"
    sudo lxc exec wails-debian-test  -- /bin/bash
    
    initTestingWailsSetup
    
    echo -e "is wails init & wails build -d enough? (y/n)"
    read input
    if [ $input = "y" ]
    then
        buildWails
    elif [ $input = "n" ]
    then
        enterCommand
    fi
}

# archlinux
archlinux(){
    echo -e "starting archlinux lxc"
    sudo lxc start wails-archlinux-test
    echo -e "exec in the bin/bash of the container"
    sudo lxc exec wails-archlinux-test  -- /bin/bash
    
    initTestingWailsSetup
    
    echo -e "is wails init & wails build -d enough? (y/n)"
    read input
    if [ $input = "y" ]
    then
        buildWails
    elif [ $input = "n" ]
    then
        enterCommand
    fi
}

initTestingWailsSetup() {
    echo -e "replacing previous wails installation with specified git/branch"

    sudo lxc exec wails-centos-test  -- rm /root/go/bin/wails && rm -rf /root/wails && cd /root/ && git clone $1 && cd /root/wails && git branch $2 && git checkout $2 && cd /root/wails/cmd/wails && go install
}

buildWails(){
    sudo lxc exec wails-centos-test  -- rm -rf /root/default-project && cd /root && wails init && wails build -d
}

enterCommand(){
    while true
    do
        echo -e "enter command to run (remeber it has to be with full paths and one liner)"
        read input
        sudo lxc exec wails-centos-test  -- $input
    done
}

# after testing part
closing(){
    echo -e "stoping all containers"
    sudo lxc stop wails-centos-test
    sudo lxc stop wails-fedora-test
    sudo lxc stop wails-archlinux-test
    sudo lxc stop wails-debian-test
}

#
# // starting the script
#
echo -e "Starting the script"
echo -e "enter https github link"
read gitUrl
echo -e "enter branch"
read newBranch
echo -e "enter either a name (ex. debian9,centos7,fedora30,archlinux) or all: "
read input
if [ $input = "all" ]
then
    debian9 $gitUrl $newBranch
    centos7 $gitUrl $newBranch
    fedora30 $gitUrl $newBranch
    archlinux $gitUrl $newBranch
    closing
    
elif [ $input = "debian9" ]
then
    debian9 $gitUrl $newBranch
    closing
    
elif [ $input = "centos7" ]
then
    centos7 $gitUrl $newBranch
    closing
    
elif [ $input = "fedora30" ]
then
    fedora30 $gitUrl $newBranch
    closing
    
elif [ $input = "archlinux" ]
then
    archlinux $gitUrl $newBranch
    closing
fi

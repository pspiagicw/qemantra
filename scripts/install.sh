#!/bin/bash

set -e pipefail

VERSION="v0.2.1"
URL="https://github.com/pspiagicw/qemantra/releases/download/${VERSION}"
INSTALL_LOCATION="$HOME/.local/bin"
GREEN='\033[0;32m'
RED='\033[1;31m'
NC='\033[0m'

header() {
    echo -e "
                                 | |            
  __ _  ___ _ __ ___   __ _ _ __ | |_ _ __ __ _ 
 / _' |/ _ \ '_ ' _ \ / _' | '_ \| __| '__/ _' |
| (_| |  __/ | | | | | (_| | | | | |_| | | (_| |
 \__, |\___|_| |_| |_|\__,_|_| |_|\__|_|  \__,_|
    | |                                         
    |_|
    Version '0.2.1' ${GREEN}
    [*] This installer will automatically detect the type of system and start installing qemantra
    [*] Installation by default be in $HOME/.local/bin/
    ${NC}
    "

}
linux_install() {
    arch=$(uname -m)

    if (( "$arch" == "x86_64 " ))
    then
        linux_install_x86_64
    elif (( "$arch" == "arm64 "))
    then
        linux_install_arm64
    elif (( "$arch" == "i386 "))
    then
        linux_install_i386
    fi
}

print_system() {
    echo -e "${GREEN}[*] Detected a ${1} system. Installing binary...${NC}"
    sleep 2

}
linux_install_x86_64() {
    print_system "x86_64"
    BINARY="qemantra-0.2.0-linux-amd64"
    download "$BINARY"
}

linux_install_arm64() {
    print_system "ARM64"
    BINARY="qemantra-0.2.0-linux-arm"
    download "$BINARY"
}

linux_install_i386() {
    print_system "i386"
    BINARY="qemantra-0.2.0-linux-i386"
    download "$BINARY"
}

# Pass install BINARY_NAME 
download() {
    BINARY=$1
    FILE="${BINARY}.tar.gz"
    FILEURL="$URL/$FILE"
    curl -L "$FILEURL" | tar xzvf - "$BINARY"
    install "${BINARY}"

}
install() {
    mv ${1} "${INSTALL_LOCATION}/qemantra"
}

main() {
    if (( "$OSTYPE" == "linux-gnu" ))
    then
        linux_install
    else
        echo "Only Linux supported for now :("
    fi
}
header
main

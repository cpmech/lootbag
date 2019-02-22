#!/bin/bash

set -e

install_and_test(){
    HERE=`pwd`
    PKG=$1
    DOTEST=$2
    echo
    echo
    echo "=== compiling $PKG ============================================================="
    cd $PKG
    touch *.go
    go install
    if [ "$DOTEST" -eq 1 ]; then
        go test
    fi
    cd $HERE
}

for pkg in check lio neto; do
    install_and_test $pkg 1
done

echo
echo "=== SUCCESS! ============================================================"

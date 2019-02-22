#!/bin/bash

FILE="*.go"

while true; do
    inotifywait -q -e modify $FILE
    echo
    echo
    echo
    echo
    echo
    echo
    #go test -run GetParamOrPanic01
    #go test -run GetParamOrPanic02
    #go test -run GetParam01
    #go test -run RedirectToHTTPS
    #go test -run CheckGET01
    #go test -run CheckGET02
    #go test -run Restrict01
    #go test -run SendRequest01
    #go test -run SendRequest02
    #go test -run SendRequest03
    #go test -run SendRequest04
    #go test -run Results01
    #go test -run Results02
    #go test -run Results03
    #go test -run SendForm01
    go test -run SendForm02
done

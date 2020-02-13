#!/bin/bash

P=$1

DIR=$(mktemp -d)

(
    cd $DIR
    export GO111MODULE=off
    export GOPATH=$PWD
    go get $P
    cd src/$P
    go build -o artifact .
    cat artifact
)

rm -rf $DIR

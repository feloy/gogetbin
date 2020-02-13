#!/bin/bash

P=$1

export GO111MODULE=off
export GOPATH=$PWD
go get $P
(
    cd src/$P
    go build -o artifact .
    cat artifact
)

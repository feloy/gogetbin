#!/bin/bash

# This script gets a Go package and builds it
# Args:
# #1: go package path
#
# Exit statuses
# 4: go get failed
# 5: go build failed

# Output:
# stdout: built binary
# stderr: messages

P=$1

DIR=$(mktemp -d)

(
    cd $DIR
    export GO111MODULE=off
    export GOPATH=$PWD
    go get $P || exit 4
    cd src/$P
    go build -o artifact . || exit 5
) 1>&2 || exit $? # stdout/stderr are merged to stderr and exit status is bubbled up if not 0

cat $DIR/src/$P/artifact # binary is output to stdout

rm -rf $DIR > /dev/null 2>&1

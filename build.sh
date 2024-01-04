#!/bin/bash

docker run --rm -i \
    -v /go:/go \
    -v ${PWD}:/workspace \
    --network host \
    golang:1.19-alpine \
    sh -c 'cd /workspace && sh publish.sh'
#!/bin/bash

# build MacOS
GOOS=darwin GOARCH=amd64 go build
zip MacOS.zip ./got && rm -rf ./got

# build Linux
GOOS=linux GOARCH=amd64 go build
zip Linux.zip ./got && rm -rf ./got


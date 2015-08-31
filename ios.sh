#!/bin/sh

#darwin
#build

#CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 ./make.bash

#go build

CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build


# 交叉编译
# go 1.5 之后，交叉编译变得很简单，只需执行以上一行即可。
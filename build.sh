#!/bin/sh
export GOPATH=/go/src
cd /go/src
cd /go/src/udptest
go build -o udpclient udpclient.go
go build -o udpserver udpserver.go



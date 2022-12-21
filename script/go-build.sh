#!/bin/bash



pwd

go build logFile.com/log-file-go

[ ! -d "./go-build" ] && mkdir ./go-build
[ ! -d "./go-build/log-file-go" ] && mkdir ./go-build/log-file-go
[ ! -f "./go-build/log-file-go/log-file-go" ] && rm -rf ./go-build/log-file-go/log-file-go 

mv ./log-file-go ./go-build/log-file-go

chmod 777 ./go-build/log-file-go/log-file-go


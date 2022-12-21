#!/bin/bash



pwd

go build logFile.com/log-file-go

[ ! -d "./go-build" ] && mkdir $HOME/go-build
[ ! -d "./go-build/log-file-go" ] && mkdir $HOME/go-build/log-file-go
[ ! -f "$HOME/go-build/log-file-go/log-file-go" ] && rm -rf $HOME/go-build/log-file-go/log-file-go 

mv ./log-file-go $HOME/go-build/log-file-go

chmod 777 $HOME/go-build/log-file-go/log-file-go


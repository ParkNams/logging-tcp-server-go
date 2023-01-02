#!/bin/bash

cd $1

pwd

echo ":$3" "$1$2.prof" 

fuser -k -n tcp porf

go tool pprof -http ":$3" "$2.prof"
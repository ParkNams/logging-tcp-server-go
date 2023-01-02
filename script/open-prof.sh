#!/bin/bash

cd $1

pwd

echo ":6061" "$1$2.prof"

fuser -k -n tcp porf

go tool pprof -http ":6061" "$2.prof"
#!/bin/bash

cd "/logging-batch-go/logs/mem"

pwd

echo ":6062" "openFile.prof" 

fuser -k -n tcp 6062

go tool pprof -http ":6062" "openFile.prof"
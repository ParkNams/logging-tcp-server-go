#!/bin/bash

cd "/logging-batch-go/logs/cpu"

pwd

echo ":6061 openFile.prof" 

fuser -k -n tcp 6061

go tool pprof -http ":6061" "openFile.prof"
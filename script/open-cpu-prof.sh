#!/bin/bash

cd "/home/ec2-user/logs/cpu"

pwd

echo ":6061 openHttp.prof" 

fuser -k -n tcp 6061

go tool pprof -http :6061 openHttp.prof
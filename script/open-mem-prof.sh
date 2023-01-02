#!/bin/bash

cd "/home/ec2-user/logs/mem"

pwd

echo ":6062" "openHttp.prof" 

fuser -k -n tcp 6062

go tool pprof -http :6062 openHttp.prof
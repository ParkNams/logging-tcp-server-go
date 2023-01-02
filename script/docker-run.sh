#!/bin/bash

docker stop logging-batch-go-container
docker rmi logging-batch-go-image
docker build -t logging-batch-go-image .

echo "$HOME" | while read path ;
do
    docker run -p 8000:8000 -p 6061:6061 -it --name logging-batch-go-container -d --rm -v $path/logs:/logging-batch-go/logs logging-batch-go-image
done
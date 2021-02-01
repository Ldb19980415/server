#!/bin/bash

docker build -t sharelife:v1 .
docker rm -f `docker ps -aq --filter name=sharelifev1`
docker run -d --name=sharelifev1 -p 3005:3005 sharelife:v1
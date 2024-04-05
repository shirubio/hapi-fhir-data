#!/bin/sh

docker compose -f docker-compose.yaml down --remove-orphans

docker stop $(docker ps -a -q)
docker rm $(docker ps -a -q)

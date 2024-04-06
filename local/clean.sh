#!/bin/sh

./stop.sh
docker volume rm postgres-data

#docker image prune -a -f
#docker system prune -a -f
#docker volume prune -f
#docker network prune -f
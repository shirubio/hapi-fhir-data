#!/bin/sh

export REDIS_PASSWORD_ON_HOST="replace-with-the-password-for-your-local-dev-server-only"

./stop.sh

docker volume rm postgres-data

docker volume create --name=postgres-data

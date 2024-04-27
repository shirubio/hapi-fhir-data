#!/bin/sh

./stop.sh
export REDIS_PASSWORD_ON_HOST="replace-with-the-password-for-your-local-dev-server-only"
docker compose -f docker-compose.yaml up -d


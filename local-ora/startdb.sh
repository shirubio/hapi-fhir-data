#!/bin/sh

./stop.sh
docker compose -f docker-compose.yaml up -d oracle-fhir


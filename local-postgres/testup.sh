#!/bin/sh
URL="http://localhost:8088/fhir/Patient/"

status_code=$(curl -o /dev/null -s -w "%{http_code}" "$URL")

if [[ "$status_code" -ge 200 && "$status_code" -lt 300 ]]; then
    echo "Server is UP"
else
    echo "Server is not ready yet, wait a few more seconds"
fi
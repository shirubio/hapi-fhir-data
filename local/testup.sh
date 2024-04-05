#!/bin/sh

# Just a simple visual indicator that the servers are up and running

curl http://localhost:8080/fhir/swagger-ui/
curl http://localhost:8080/fhir/Patient/


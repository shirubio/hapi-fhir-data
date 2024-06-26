#!/bin/sh

# Just a simple visual indicator that the servers are up and running

curl http://localhost:8080/fhir/Patient?_summary=true
curl http://localhost:8080/fhir/Observation?_summary=true

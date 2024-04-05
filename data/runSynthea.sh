#!/bin/sh

# https://github.com/synthetichealth/synthea/wiki/Basic-Setup-and-Running

numPatients=${1:-101}
echo "Generating $numPatients patients"

#Main
rm -rf output
java -jar synthea-with-dependencies.jar -p "$numPatients" Illinois Chicago

mkdir output/dependencies
mv output/fhir/practitionerInformation*.json output/dependencies
mv output/fhir/hospitalInformation*.json output/dependencies

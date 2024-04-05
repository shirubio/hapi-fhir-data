PGHD FHIR Federator Service - Local Dev Environment
===============================================

How to create a local dev environment for HAPI FHIR Service.:

Pre-requisites:
- Docker 
- Docker Compose
- Git
- Java 11+
- Maven 3.6+
- Go 1.18+

Optional. Clone Hapi FHIR JPA Server Starter project:
   https://github.com/hapifhir/hapi-fhir-jpaserver-starter



1. Run your local HAPI FHIR server on port 8080.  

2. Generate test data.

   1. Download Synthea from GitHub: 
https://github.com/synthetichealth/synthea/releases/download/master-branch-latest/synthea-with-dependencies.jar
   2. run:
   ./runSynthea.sh [Number of Patients]     Default is 101 patients.

3. Load the data into local hapi fhir servers. 
   - Go to the 'data' folder and run:
   go run loadTestData.go 
    
   This can be very slow if you asked for a large number of patients.
      
   After loading the data, you can delete all subfolders in the 'data' folder. 

4. Hapi coding!
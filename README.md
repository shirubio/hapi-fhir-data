Create test data for HAPI FHIR Service - Local Dev Environment
=

Pre-requisites:
- Docker (Docker Desktop for Mac or Windows) 
- Docker Compose
- Git
- Java 17+
- Maven 3.0+
- Go 1.20+

1. Preparing Test Data
   1. Go to the 'data' folder 
   2. Download Synthea from GitHub:
         https://github.com/synthetichealth/synthea/releases/download/master-branch-latest/synthea-with-dependencies.jar
   3. run:
      ./runSynthea.sh [Number of Patients]     Default is 101 patients.
   4. You should look at all the options in Synthea. There are many options to generate test data.

2. Load the data into local hapi fhir servers. 
   - Go to the 'data' folder and run:
   go run loadTestData.go 
    
   This can be very slow if you asked for a large number of patients.
      
   After loading the data, you can delete all folders in the 'data' folder. 

4. Hapi coding!
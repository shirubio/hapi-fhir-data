package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func main() {
	fmt.Println("************************* Version - 2024-04-05")
	loadTestData(0, "./output/dependencies/")
	loadTestData(0, "./output/fhir")
}

// You can populate multiple service URLs here
var serviceURLs = []string{
	"http://localhost:8080/fhir",
	//"http://localhost:8081/fhir",
	//"http://localhost:8082/fhir",
}

func loadTestData(urlIndex int, directoryPath string) {
	fileInfos, err := os.ReadDir(directoryPath)
	if err != nil {
		log.Fatal("Error reading payload directory:", err)
	}

	sem := make(chan struct{}, 5)

	count := 0
	for _, fileInfo := range fileInfos {
		filePath := filepath.Join(directoryPath, fileInfo.Name())
		count++
		if filepath.Ext(filePath) != ".json" {
			continue
		}

		fmt.Printf("Ready to process file %s - %d out of %d\n", filePath, count, len(fileInfos))

		payload, err := os.ReadFile(filePath)
		if err != nil {
			log.Printf("Error reading payload file %s: %s", filePath, err)
			continue
		}

		sem <- struct{}{}

		go func(payload []byte, filePath string, url string) {
			defer func() { <-sem }()
			processFileWithRetry(payload, filePath, url, 5)
		}(payload, filePath, serviceURLs[urlIndex])
	}

	// Wait for all goroutines to finish
	for i := 0; i < cap(sem); i++ {
		sem <- struct{}{}
	}
}

func processFileWithRetry(payload []byte, filePath string, url string, retries int) {
	for retry := 0; retry < retries; retry++ {
		// Create a new HTTP POST request
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
		if err != nil {
			log.Printf("\t*** Failed - Retrying - Error creating request for file %s and URL %s: %s", filePath, url, err)
			time.Sleep(2 * time.Second)
			continue
		}

		// Tell the server that we expect a JSON response
		req.Header.Set("Accept", "application/fhir+json")
		req.Header.Set("Content-Type", "application/fhir+json")

		// Send the HTTP request
		client := &http.Client{}
		fmt.Printf("Calling FHIR for file %s on server %s\n", filePath, url)
		resp, err := client.Do(req)
		if err != nil {
			log.Printf("\t*** Failed - Retrying - Error sending request for file %s and URL %s: %s", filePath, url, err)
			time.Sleep(2 * time.Second)
			continue
		}
		fmt.Printf("\t*** Completed processing file %s on server: %s with Status: %s\n", filePath, url, resp.Status)

		if resp.StatusCode >= 200 && resp.StatusCode < 300 {
			renameFile(filePath)
			break
		}
		fmt.Printf("\t*** Failed - Retrying ***\n")
		time.Sleep(10 * time.Second)
	}
}

func renameFile(path string) {
	// Rename the file to have a .done extension, this allow to stop and restart the process without reprocessing the same files
	newPath := path + ".done"
	_ = os.Rename(path, newPath)
	fmt.Printf("\t*** Renamed file %s to %s\n", path, newPath)
}

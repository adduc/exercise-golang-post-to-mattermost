package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

/*
  Goal:
    - Post a simple "Hello, World!" message to a mattermost incoming
	  webhook

  Steps:
	- load .env file (if it exists)
	- check for required environment variables
	- create a new client
	- post a message

  Constraints:
    - third-party libraries should be avoided (dotenv is allowed)

  Environment Variables:
    - WEBHOOK_URL: the URL of the incoming webhook

  Links:
    - https://developers.mattermost.com/integrate/webhooks/incoming/#parameters
*/

// loadEnv loads the .env file if it exists
func loadEnv() {
	err := godotenv.Load()
	if err != nil && !os.IsNotExist(err) {
		// only log the error if it's not a "file not found" error
		log.Fatal(err)
	}
}

// checkEnv checks for required environment variables
func checkEnv() {
	webhookURL := os.Getenv("WEBHOOK_URL")
	if webhookURL == "" {
		log.Fatal("WEBHOOK_URL environment variable is required")
	}
}

// createHttpClient creates a new http client
func createHttpClient() *http.Client {
	client := &http.Client{}

	return client
}

// postMessage posts a message to the mattermost incoming webhook
func postMessage(client *http.Client, webhookURL string) {
	// create a new request
	req, err := http.NewRequest("POST", webhookURL, nil)
	if err != nil {
		log.Fatal(err)
	}

	// set the content type
	req.Header.Set("Content-Type", "application/json")

	// define the payload struct
	type Payload struct {
		Text string `json:"text"`
	}

	// create the payload
	payload := Payload{
		Text: "Hello, World!",
	}

	// marshal the payload to JSON
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err)
	}

	// set the request body to the JSON payload
	req.Body = io.NopCloser(strings.NewReader(string(payloadBytes)))

	// send the request
	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	// close the response body
	defer resp.Body.Close()
}

func main() {
	loadEnv()
	checkEnv()

	client := createHttpClient()

	webhookUrl := os.Getenv("WEBHOOK_URL")
	postMessage(client, webhookUrl)
}

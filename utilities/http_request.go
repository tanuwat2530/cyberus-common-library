package utilities

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Common function to make HTTP requests
func makeRequest(method, url string, jsonData []byte) (string, error) {
	client := &http.Client{}

	var req *http.Request
	var err error

	if method == http.MethodGet {
		req, err = http.NewRequest(method, url, nil)
	} else {
		req, err = http.NewRequest(method, url, bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")
	}

	if err != nil {
		return "", fmt.Errorf("error creating request: %w", err)
	}

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	// Read response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response: %w", err)
	}

	return string(body), nil
}

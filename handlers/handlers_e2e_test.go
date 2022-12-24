package handlers_test

import (
	"net/http"
	"testing"
)

const (
	ENDPOINT_PING = "http://localhost:9090/ping"
)

func TestApplicationLiveness(t *testing.T) {
	// Create a new request using http to get the file
	req, err := http.NewRequest("GET", ENDPOINT_PING, nil)
	if err != nil {
		t.Errorf("TestApplicationLiveness: Cannot create new request: %s", err.Error())
		return
	}

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("TestApplicationLiveness: Cannot send request: %s", err.Error())
		return
	}

	// Check the response
	if resp.StatusCode != http.StatusOK {
		t.Errorf("TestApplicationLiveness resp.StatusCode = %v, want %v", resp.StatusCode, http.StatusOK)
		return
	}
}

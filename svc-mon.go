package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

// Service represents a service to monitor
type Service struct {
	Name          string        `json:"name"`
	URL           string        `json:"url"`
	IP            string        `json:"ip"`
	Port          string        `json:"port"`
	LocalImage    string        `json:"localImage"`
	CheckInterval time.Duration `json:"checkInterval"`
}

// Config represents the configuration for the program
type Config struct {
	Services []Service `json:"services"`
}

// HTTPClient represents an interface for the HTTP client to enable testing
type HTTPClient interface {
	Get(url string) (*http.Response, error)
}

// DefaultHTTPClient represents the default implementation of HTTPClient using the actual http.Client
type DefaultHTTPClient struct{}

// Get implements the Get method for DefaultHTTPClient
func (c *DefaultHTTPClient) Get(url string) (*http.Response, error) {
	return http.Get(url)
}

// Monitor continuously monitors the services and reports their availability
func svcMonitor(services []Service, client HTTPClient) {
	for _, service := range services {
		go func(service Service) {
			for {
				status := "up"
				resp, err := client.Get(service.URL)
				if err != nil || resp.StatusCode != http.StatusOK {
					status = "down"
				}
				fmt.Printf("[%s] %s is %s\n", time.Now().Format(time.RFC1123), service.Name, status)
				time.Sleep(service.CheckInterval * time.Second)
			}
		}(service)
	}
}

func TestSvcMonitor(t *testing.T) {
	// Create a test service
	testService := Service{
		Name:          "TestService",
		URL:           "http://example.com",
		IP:            "127.0.0.1",
		Port:          "8080",
		LocalImage:    "local-image",
		CheckInterval: 1,
	}

	// Create a mock HTTP client
	mockClient := &MockHTTPClient{
		Resp: &http.Response{
			StatusCode: http.StatusOK,
		},
		Err: nil,
	}

	// Call the Monitor function with the test service and mock HTTP client
	go svcMonitor([]Service{testService}, mockClient)

	// Sleep for a short duration to allow Monitor to run
	time.Sleep(3 * time.Second)
}

// MockHTTPClient is a mock implementation of HTTPClient for testing
type MockHTTPClient struct {
	Resp *http.Response
	Err  error
}

// Get implements the Get method for MockHTTPClient
func (m *MockHTTPClient) Get(url string) (*http.Response, error) {
	return m.Resp, m.Err
}

func TestMain(t *testing.T) {
	// You can write additional test cases for the main function if needed.
	// For testing the main function, consider using TestMain(m *testing.M) and os.Exit to capture the behavior.
}

// The main function should be left empty for testing, as the actual functionality is in the Monitor function.
func main() {}

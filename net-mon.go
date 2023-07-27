package main

import (
	"fmt"
	"net"
	"testing"
	"time"
)

// NetworkChecker represents an interface for the network connectivity checker to enable testing
type NetworkChecker interface {
	Dial(network, address string) (net.Conn, error)
}

// DefaultNetworkChecker represents the default implementation of NetworkChecker using net.Dial
type DefaultNetworkChecker struct{}

// Dial implements the Dial method for DefaultNetworkChecker
func (c *DefaultNetworkChecker) Dial(network, address string) (net.Conn, error) {
	return net.Dial(network, address)
}

// Monitor continuously monitors the network connectivity and reports its status
func netMonitor(interval time.Duration, host string, checker NetworkChecker) {
	for {
		conn, err := checker.Dial("tcp", host)
		status := "up"
		if err != nil {
			status = "down"
		} else {
			conn.Close()
		}
		fmt.Printf("[%s] Network is %s\n", time.Now().Format(time.RFC1123), status)
		time.Sleep(interval)
	}
}

func TestNetMonitor(t *testing.T) {
	// Create a test host and interval
	testHost := "example.com:80"
	testInterval := 1 * time.Second

	// Create a mock network checker
	mockChecker := &MockNetworkChecker{
		Err: nil,
	}

	// Call the netMonitor function with the test host, interval, and mock checker
	go netMonitor(testInterval, testHost, mockChecker)

	// Sleep for a short duration to allow netMonitor to run
	time.Sleep(3 * time.Second)
}

// MockNetworkChecker is a mock implementation of NetworkChecker for testing
type MockNetworkChecker struct {
	Err error
}

// Dial implements the Dial method for MockNetworkChecker
func (m *MockNetworkChecker) Dial(network, address string) (net.Conn, error) {
	return nil, m.Err
}

func TestMain(t *testing.T) {
	// You can write additional test cases for the main function if needed.
	// For testing the main function, consider using TestMain(m *testing.M) and os.Exit to capture the behavior.
}

// The main function should be left empty for testing, as the actual functionality is in the netMonitor function.
func main() {}

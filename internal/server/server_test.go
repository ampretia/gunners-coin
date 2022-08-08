//go:build !integration

// Package server_test contains simple unit and integration tests for exported
// stuff in server.
package server_test

import (
	"testing"
)

func TestServer(t *testing.T) {
	t.Run("NewServer", testNewServer)
}

// Simple test case structure
func testNewServer(t *testing.T) {

}

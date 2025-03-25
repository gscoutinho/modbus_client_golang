package client

import (
	"net"
	"testing"
	"time"
)

// TestTCPClientExecute tests the Execute method of TCPClient by setting up
// a temporary server that simulates a Modbus TCP device.
func TestTCPClientExecute(t *testing.T) {
	// Start a temporary TCP server on localhost with an available port.
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("failed to listen: %v", err)
	}
	defer ln.Close()

	// The expected response is built as follows:
	// MBAP header:
	//   Transaction ID: 0x00, 0x02
	//   Protocol ID:    0x00, 0x00
	//   Length:         0x00, 0x05  (1 byte for Unit ID + 4 bytes for PDU)
	//   Unit Identifier: 0x01
	// PDU:
	//   Data bytes: 0xAA, 0xBB, 0xCC, 0xDD
	expectedResponse := []byte{
		0x00, 0x02, // Transaction ID
		0x00, 0x00, // Protocol ID
		0x00, 0x05, // Length
		0x01,                   // Unit Identifier
		0xAA, 0xBB, 0xCC, 0xDD, // PDU bytes
	}

	// Start a goroutine to simulate the Modbus TCP server.
	done := make(chan struct{})
	go func() {
		defer close(done)
		conn, err := ln.Accept()
		if err != nil {
			t.Errorf("server failed to accept connection: %v", err)
			return
		}
		defer conn.Close()

		// Optionally, read the incoming request (not used in this test).
		reqBuf := make([]byte, 1024)
		_, _ = conn.Read(reqBuf)

		// Write the simulated response.
		_, err = conn.Write(expectedResponse)
		if err != nil {
			t.Errorf("server failed to write response: %v", err)
		}
	}()

	// Get the actual port from the listener.
	addr := ln.Addr().(*net.TCPAddr)
	port := addr.Port

	// Create a TCPClient that will connect to our temporary server.
	client := NewTCPClient("127.0.0.1", 5*time.Second, port, nil)

	// Create a dummy request (the content isn't important for this test).
	dummyRequest := []byte{0x01, 0x03, 0x00, 0x00, 0x00, 0x02}

	// Execute the request.
	response, err := client.Execute(dummyRequest)
	if err != nil {
		t.Fatalf("Execute() error: %v", err)
	}

	// Verify that the response matches our expected response.
	if len(response) != len(expectedResponse) {
		t.Fatalf("response length mismatch: expected %d, got %d", len(expectedResponse), len(response))
	}

	for i, b := range expectedResponse {
		if response[i] != b {
			t.Errorf("response mismatch at index %d: expected 0x%02X, got 0x%02X", i, b, response[i])
		}
	}

	// Wait for the server goroutine to finish.
	<-done
}

package client

import (
	"net"
	"testing"
	"time"
)

// startTestServer starts a simple TCP server on localhost for testing purposes.
// It listens on a free port and accepts connections until it is closed.
func startTestServer(t *testing.T) net.Listener {
	ln, err := net.Listen("tcp", "127.0.0.1:502")
	if err != nil {
		t.Fatalf("failed to start test server: %v", err)
	}
	go func() {
		for {
			conn, err := ln.Accept()
			if err != nil {
				return // Likely closed
			}
			// For testing, we just wait a bit then close the connection.
			time.Sleep(100 * time.Millisecond)
			conn.Close()
		}
	}()
	return ln
}

func TestTCPConnectionPool_GetPut(t *testing.T) {
	ln := startTestServer(t)
	defer ln.Close()

	addr := ln.Addr().String()

	// Create a pool with capacity 2.
	pool := NewTCPConnectionPool(addr, 2*time.Second, 2)

	// Get a connection from the pool.
	conn1, err := pool.Get()
	if err != nil {
		t.Fatalf("failed to get connection: %v", err)
	}
	if conn1 == nil {
		t.Fatalf("expected a valid connection, got nil")
	}
	// Return the connection to the pool.
	pool.Put(conn1)

	// Get again: should return the same connection (or one from the pool).
	conn2, err := pool.Get()
	if err != nil {
		t.Fatalf("failed to get connection second time: %v", err)
	}
	// Depending on the pool implementation, conn1 and conn2 may or may not be the same.
	pool.Put(conn2)

	// Now test capacity enforcement.
	// Fill the pool.
	connA, _ := pool.Get()
	connB, _ := pool.Get()
	pool.Put(connA)
	pool.Put(connB)

	// Create an extra connection manually.
	connC, err := net.DialTimeout("tcp", addr, 2*time.Second)
	if err != nil {
		t.Fatalf("failed to dial extra connection: %v", err)
	}
	// Put the extra connection into the pool. Since the pool is full, it should close connC.
	pool.Put(connC)

	// Try writing to connC. Since it should be closed, the write should fail.
	_, err = connC.Write([]byte("test"))
	if err == nil {
		t.Errorf("expected extra connection to be closed, but write succeeded")
	}
}

func TestTCPConnectionPool_Close(t *testing.T) {
	ln := startTestServer(t)
	defer ln.Close()

	addr := ln.Addr().String()
	pool := NewTCPConnectionPool(addr, 2*time.Second, 2)

	// Get a connection and return it so that the pool is not empty.
	conn, err := pool.Get()
	if err != nil {
		t.Fatalf("failed to get connection: %v", err)
	}
	pool.Put(conn)

	// Close the pool.
	pool.Close()

	// After closing, the pool channel should be drained.
	for conn := range pool.pool {
		if conn != nil {
			t.Errorf("expected no connection after pool close, but got one")
		}
	}

	// Attempt to put a new connection into the closed pool.
	newConn, err := net.DialTimeout("tcp", addr, 2*time.Second)
	if err != nil {
		t.Fatalf("failed to dial new connection: %v", err)
	}
	pool.Put(newConn)
	// Verify that newConn is closed by trying to write to it.
	_, err = newConn.Write([]byte("test"))
	if err == nil {
		t.Errorf("expected new connection to be closed after pool.Put on a closed pool")
	}
}

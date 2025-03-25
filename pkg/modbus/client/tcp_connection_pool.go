package client

import (
	"fmt"
	"net"
	"sync"
	"time"
)

type TCPConnectionPool struct {
	address string
	timeout time.Duration
	pool    chan net.Conn
	mu      sync.Mutex
	closed  bool
}

func NewTCPConnectionPool(address string, timeout time.Duration, maxConnections int) *TCPConnectionPool {
	return &TCPConnectionPool{
		address: address,
		timeout: timeout,
		pool:    make(chan net.Conn, maxConnections),
		closed:  false,
	}
}

func (p *TCPConnectionPool) Get() (net.Conn, error) {
	select {
	case conn := <-p.pool:
		return conn, nil
	default:
		conn, err := net.DialTimeout("tcp", p.address, p.timeout)
		if err != nil {
			return nil, fmt.Errorf("failed to establish connection to %s: %w", p.address, err)
		}
		return conn, nil
	}
}

func (p *TCPConnectionPool) Put(conn net.Conn) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.closed {
		conn.Close()
		return
	}

	select {
	case p.pool <- conn:
		//conn returned to the pool
	default:
		conn.Close()
	}
}

func (p *TCPConnectionPool) Close() {
	p.mu.Lock()

	if !p.closed {
		p.closed = true
		close(p.pool)
	}

	p.mu.Unlock()

	for conn := range p.pool {
		conn.Close()
	}
}

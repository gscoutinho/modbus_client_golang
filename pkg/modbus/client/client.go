package client

import (
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

type TCPClient struct {
	Host    string
	Timeout time.Duration
	Port    int
	Pool    *TCPConnectionPool
}

func NewTCPClient(host string, timeout time.Duration, port int, pool *TCPConnectionPool) *TCPClient {
	return &TCPClient{
		Host:    host,
		Timeout: timeout,
		Port:    port,
		Pool:    pool,
	}
}

func (c *TCPClient) Execute(tcpRequest []byte) ([]byte, error) {

	//conenction establishment
	var conn net.Conn
	var err error

	address := fmt.Sprintf("%s:%d", c.Host, c.Port)

	if c.Pool != nil {
		conn, err = c.Pool.Get()
		if err != nil {
			return nil, fmt.Errorf("failed to get connection from pool: %w", err)
		}

		defer c.Pool.Put(conn)
	} else {
		conn, err = net.DialTimeout("tcp", address, c.Timeout)
		if err != nil {
			return nil, fmt.Errorf("failed to connect to %s: %w", address, err)
		}
		defer conn.Close()
	}

	conn.SetDeadline(time.Now().Add(c.Timeout))

	_, err = conn.Write(tcpRequest)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}

	header := make([]byte, 7)
	n, err := conn.Read(header)

	if err != nil {
		return nil, fmt.Errorf("failed to read Modbus header: %w", err)
	}

	if n != 7 {
		return nil, fmt.Errorf("incomplete Modbus header: expected 7 bytes, got %d", n)
	}

	length := binary.BigEndian.Uint16(header[4:6])

	remaining := int(length) - 1

	pduResp := make([]byte, remaining)

	totalRead := 0

	for totalRead < remaining {
		n, err = conn.Read(pduResp[totalRead:])
		if err != nil {
			return nil, fmt.Errorf("failed to read PDU: %w", err)
		}
		totalRead += n
	}

	response := append(header, pduResp...)
	return response, nil
}

package modbus

import (
	"fmt"
	"net"
	"time"
)

type TCPClient struct {
	Host    string
	Timeout time.Duration
	Port    int
	Pool    *any //TCPConnectionPool //optional
}

func NewTCPClient(host string, timeout time.Duration, port int, pool *any) *TCPClient {
	return &TCPClient{
		Host:    host,
		Timeout: timeout,
		Port:    port,
		Pool:    nil,
	}
}

func (c *TCPClient) Execute(request []byte) ([]byte, error) {
	var conn net.Conn
	var err error
	address := fmt.Sprintf("%s:%d", c.Host, c.Port)
	if c.Pool != nil {
		//when pool is available
		// conn, err == c.Pool.Get()
		// if err != nil{
		// 	return nil, fmt.Errorf("failed to get connection from pool: %w", err)
		// }

		// defer c.Pool.Put(conn)
	} else {
		conn, err = net.DialTimeout("tcp", address, c.Timeout)
		if err != nil {
			return nil, fmt.Errorf("failed to connect to %s: %w")
		}
		defer conn.Close()
	}

	conn.SetDeadline(time.Now().Add(c.Timeout))

	_, err = conn.Write(request)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}

	return nil, nil
}

package modbus_client

import (
	"modbus_client/pkg/modbus/client"
	"time"
)

// this is the edge layer between CLI and modbus package
type ModbusClient struct {
	TCPClient *client.TCPClient
}

func NewModbusClient(host string, port int, timeout time.Duration, pool *client.TCPConnectionPool) *ModbusClient {
	return &ModbusClient{
		TCPClient: client.NewTCPClient(host, timeout, port, pool),
	}
}

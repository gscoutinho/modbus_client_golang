package modbus

type ModbusHeader struct {
	FC          FunctionCode
	SlaveID     byte
	DataAddress [2]byte
}

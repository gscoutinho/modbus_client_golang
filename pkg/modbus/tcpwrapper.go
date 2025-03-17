package modbus

type TCPRequestWrapper struct {
	TransactionID uint16
	ProtocolID    uint16
	MessageLength uint16
	ModbusFrame   []byte
}

func (wrap *TCPRequestWrapper) Build() ([]byte, error) {
	frame := make([]byte, 6+len(wrap.ModbusFrame))

}

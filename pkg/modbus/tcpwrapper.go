package modbus

import (
	"encoding/binary"
)

type TCPRequestWrapper struct {
	TransactionID uint16
	ProtocolID    uint16
	MessageLength uint16
	ModbusFrame   []byte
}

func (wrap *TCPRequestWrapper) Build() ([]byte, error) {
	frame := make([]byte, 6+len(wrap.ModbusFrame))

	//modbus tcp frame:
	// transcation ID 0001
	// protocol ID 0000
	// message length 00XX (bytes to folllow)
	// modbus frame
	wrap.MessageLength = uint16(len(wrap.ModbusFrame))

	frame[0] = 0b0000
	frame[1] = 0b0001
	frame[2] = 0b0000
	frame[3] = 0b0000
	binary.BigEndian.PutUint16(frame[4:6], wrap.MessageLength)
	copy(frame[6:], wrap.ModbusFrame)

	return frame, nil
}

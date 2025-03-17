package modbus

import (
	"fmt"
)

type ReadingResponse struct {
	Header    ModbusHeader
	ByteCount uint16
	Response  []byte
}

type SingleWritingResponse struct {
	Header       ModbusHeader
	ValueWritten []byte
}

type MultipleWritingResponse struct {
	Header          ModbusHeader
	QuantityWritten []byte
}

func (r *ReadingResponse) Build() ([]byte, error) {

	if r.ByteCount != uint16(len(r.Response)) {
		return nil, fmt.Errorf("ByteCount mismatch: expected %d, got %d", r.ByteCount, len(r.Response))
	}

	// Frame layout
	// [0] SlaveID
	// [1] FunctionCode
	// [2] Qty of data bytes to follow
	// [n] values
	frame := make([]byte, 3+r.ByteCount)
	frame[0] = r.Header.SlaveID
	frame[1] = byte(r.Header.FC)
	frame[2] = byte(r.ByteCount)

	for i := 0; i < len(r.Response); i++ {
		frame[i+3] = r.Response[i]
	}

	return frame, nil
}

func (r *SingleWritingResponse) Build() ([]byte, error) {

	// Frame layout
	// [0] SlaveID
	// [1] FunctionCode
	// [2] DataAddress of first coil/stat/reg: Address High
	// [3] DataAddress of first coil/stat/reg: Address Low
	// [n] written value

	frame := make([]byte, 4+len(r.ValueWritten))
	frame[0] = r.Header.SlaveID
	frame[1] = byte(r.Header.FC)
	frame[2] = r.Header.DataAddress[0]
	frame[3] = r.Header.DataAddress[1]

	for i := 0; i < len(r.ValueWritten); i++ {
		frame[i+4] = r.ValueWritten[i]
	}

	return frame, nil
}

func (r *MultipleWritingResponse) Build() ([]byte, error) {
	//frame layout
	// [0] SlaveID
	// [1] FunctionCode
	// [2] DataAddress of first coil/stat/reg: Address High
	// [3] DataAddress of first coil/stat/reg: Address Low
	// [4] qty of written coil/stat/reg: high
	// [5] qty of written coil/stat/reg: low

	frame := make([]byte, 6)
	frame[0] = r.Header.SlaveID
	frame[1] = byte(r.Header.FC)
	frame[2] = r.Header.DataAddress[0]
	frame[3] = r.Header.DataAddress[1]
	frame[4] = r.QuantityWritten[0]
	frame[5] = r.QuantityWritten[1]

	return frame, nil
}

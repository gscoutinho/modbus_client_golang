package modbus

type ReadingRequest struct {
	Header   ModbusHeader
	Quantity uint16
}

type SingleWritingRequest struct {
	Header      ModbusHeader
	Value2Write uint16
}

type MultipleWritingRequest struct {
	Header       ModbusHeader
	Quantity     uint16
	Values2Write []byte
}

func (r *ReadingRequest) Build() ([]byte, error) {
	// Frame layout (Modbus PDU):
	//   [0] SlaveID
	//   [1] FunctionCode
	//   [2] Address High
	//   [3] Address Low
	//   [4] Quantity High
	//   [5] Quantity Low

	frame := make([]byte, 6)
	frame[0] = r.Header.SlaveID
	frame[1] = byte(r.Header.FC)
	frame[2] = r.Header.DataAddress[0]
	frame[3] = r.Header.DataAddress[1]
	frame[4] = byte(r.Quantity >> 8)
	frame[5] = byte(r.Quantity & 0xFF)

	return frame, nil
}

func (r *SingleWritingRequest) Build() ([]byte, error) {
	// Frame layout (Modbus PDU):
	//   [0] SlaveID
	//   [1] FunctionCode
	//   [2] Address High
	//   [3] Address Low
	//   [4] Value High
	//   [5] Value Low
	frame := make([]byte, 6)

	frame[0] = r.Header.SlaveID
	frame[1] = byte(r.Header.FC)
	frame[2] = r.Header.DataAddress[0]
	frame[3] = r.Header.DataAddress[1]
	frame[4] = byte(r.Value2Write >> 8)
	frame[5] = byte(r.Value2Write & 0xFF)

	return frame, nil
}

func (r *MultipleWritingRequest) Build() ([]byte, error) {

	// Build the frame (Modbus PDU for multiple write):
	// Layout:
	//   [0] SlaveID
	//   [1] FunctionCode
	//   [2-3] DataAddress (high, low)
	//   [4-5] Quantity (number of registers)
	//   [6]   Byte Count (number of data bytes)
	//   [7...] Data payload
	frame := make([]byte, 7+len(r.Values2Write))
	frame[0] = r.Header.SlaveID
	frame[1] = byte(r.Header.FC)
	frame[2] = r.Header.DataAddress[0]
	frame[3] = r.Header.DataAddress[1]
	frame[4] = byte(r.Quantity >> 8)
	frame[5] = byte(r.Quantity & 0xFF)
	frame[6] = byte(len(r.Values2Write))
	copy(frame[7:], r.Values2Write)

	return frame, nil
}

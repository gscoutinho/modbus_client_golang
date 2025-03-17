package modbus

import (
	"fmt"
	"reflect"
	"testing"
)

// Test for ReadingRequest.Build()
func TestReadingRequestBuild(t *testing.T) {
	// For this test, assume FunctionCode 0x03 represents a "read holding registers" request.
	req := &ReadingRequest{
		Header: ModbusHeader{
			FC:          0x03, // For example, FCReadHoldingRegisters = 0x03
			SlaveID:     0x01,
			DataAddress: [2]byte{0x00, 0x10},
		},
		Quantity: 0x0002,
	}

	expected := []byte{0x01, 0x03, 0x00, 0x10, 0x00, 0x02}
	frame, err := req.Build()
	if err != nil {
		t.Fatalf("ReadingRequest Build() error: %v", err)
	}
	if !reflect.DeepEqual(frame, expected) {
		t.Errorf("ReadingRequest Build() failed.\nExpected: %v\nGot:      %v", expected, frame)
	}
}

// Test for SingleWritingRequest.Build()
func TestSingleWritingRequestBuild(t *testing.T) {
	// For this test, assume FunctionCode 0x06 represents a "write single register" request.
	req := &SingleWritingRequest{
		Header: ModbusHeader{
			FC:          0x06, // For example, FCWriteSingleRegister = 0x06
			SlaveID:     0x01,
			DataAddress: [2]byte{0x00, 0x20},
		},
		Value2Write: 0x0123,
	}

	expected := []byte{0x01, 0x06, 0x00, 0x20, 0x01, 0x23}
	frame, err := req.Build()
	if err != nil {
		t.Fatalf("SingleWritingRequest Build() error: %v", err)
	}
	if !reflect.DeepEqual(frame, expected) {
		t.Errorf("SingleWritingRequest Build() failed.\nExpected: %v\nGot:      %v", expected, frame)
	}
}

// A simple converter for uint16 values.
// It converts a uint16 value into its 2-byte big-endian representation.
func Uint16Converter(v any) ([]byte, error) {
	val, ok := v.(uint16)
	if !ok {
		return nil, fmt.Errorf("value %v is not a uint16", v)
	}
	b := []byte{byte(val >> 8), byte(val & 0xFF)}
	return b, nil
}

// Test for MultipleWritingRequest.BuildWithConverter()
func TestMultipleWritingRequestBuild(t *testing.T) {
	// For this test, assume FunctionCode 0x10 represents a "write multiple registers" request.
	req := &MultipleWritingRequest{
		Header: ModbusHeader{
			FC:          0x10, // FCWriteMultipleRegisters = 0x10
			SlaveID:     0x01,
			DataAddress: [2]byte{0x00, 0x30},
		},
		Quantity:     0x0002, // Expecting 2 registers (4 bytes)
		Values2Write: []byte{0x0A, 0x0B, 0x0C, 0x0D},
	}

	// Expected frame breakdown:
	// [0] SlaveID:          0x01
	// [1] FunctionCode:     0x10
	// [2-3] DataAddress:    0x00, 0x30
	// [4-5] Quantity:       0x00, 0x02
	// [6] Byte Count:       0x04 (4 bytes of payload)
	// [7-10] Payload:       0x0A, 0x0B, 0x0C, 0x0D
	expected := []byte{0x01, 0x10, 0x00, 0x30, 0x00, 0x02, 0x04, 0x0A, 0x0B, 0x0C, 0x0D}

	frame, err := req.Build()
	if err != nil {
		t.Fatalf("MultipleWritingRequest Build() error: %v", err)
	}
	if !reflect.DeepEqual(frame, expected) {
		t.Errorf("MultipleWritingRequest Build() failed.\nExpected: %v\nGot:      %v", expected, frame)
	}
}

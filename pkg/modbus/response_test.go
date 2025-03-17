package modbus

import (
	"reflect"
	"testing"
)

func TestReadingResponseBuild_Success(t *testing.T) {
	// Construct a ReadingResponse with matching ByteCount and Response length.
	rr := &ReadingResponse{
		Header: ModbusHeader{
			FC:          0x03, // e.g. Read Holding Registers
			SlaveID:     0x01,
			DataAddress: [2]byte{0x00, 0x10},
		},
		ByteCount: 4,
		Response:  []byte{0x12, 0x34, 0x56, 0x78},
	}

	// Expected frame layout:
	// [0] SlaveID: 0x01
	// [1] FunctionCode: 0x03
	// [2] ByteCount: 0x04
	// [3..6] Payload: 0x12, 0x34, 0x56, 0x78
	expected := []byte{0x01, 0x03, 0x04, 0x12, 0x34, 0x56, 0x78}

	frame, err := rr.Build()
	if err != nil {
		t.Fatalf("ReadingResponse Build() returned error: %v", err)
	}
	if !reflect.DeepEqual(frame, expected) {
		t.Errorf("ReadingResponse Build() failed.\nExpected: %v\nGot:      %v", expected, frame)
	}
}

func TestReadingResponseBuild_ByteCountMismatch(t *testing.T) {
	// Construct a ReadingResponse with mismatched ByteCount vs. Response length.
	rr := &ReadingResponse{
		Header: ModbusHeader{
			FC:          0x03,
			SlaveID:     0x01,
			DataAddress: [2]byte{0x00, 0x10},
		},
		ByteCount: 3, // Incorrect: actual response has 4 bytes.
		Response:  []byte{0x12, 0x34, 0x56, 0x78},
	}

	_, err := rr.Build()
	if err == nil {
		t.Fatal("Expected error due to ByteCount mismatch, but got nil")
	}
}

func TestSingleWritingResponseBuild(t *testing.T) {
	// Construct a SingleWritingResponse.
	swr := &SingleWritingResponse{
		Header: ModbusHeader{
			FC:          0x06, // e.g. Write Single Register
			SlaveID:     0x02,
			DataAddress: [2]byte{0x00, 0x20},
		},
		ValueWritten: []byte{0xAB, 0xCD},
	}

	// Expected frame layout:
	// [0] SlaveID: 0x02
	// [1] FunctionCode: 0x06
	// [2] DataAddress high: 0x00
	// [3] DataAddress low:  0x20
	// [4..5] Written value: 0xAB, 0xCD
	expected := []byte{0x02, 0x06, 0x00, 0x20, 0xAB, 0xCD}

	frame, err := swr.Build()
	if err != nil {
		t.Fatalf("SingleWritingResponse Build() returned error: %v", err)
	}
	if !reflect.DeepEqual(frame, expected) {
		t.Errorf("SingleWritingResponse Build() failed.\nExpected: %v\nGot:      %v", expected, frame)
	}
}

func TestMultipleWritingResponseBuild(t *testing.T) {
	// Construct a MultipleWritingResponse.
	mwr := &MultipleWritingResponse{
		Header: ModbusHeader{
			FC:          0x10, // e.g. Write Multiple Registers
			SlaveID:     0x03,
			DataAddress: [2]byte{0x00, 0x40},
		},
		QuantityWritten: []byte{0x00, 0x02}, // Representing quantity as two bytes: high, low.
	}

	// Expected frame layout:
	// [0] SlaveID: 0x03
	// [1] FunctionCode: 0x10
	// [2] DataAddress high: 0x00
	// [3] DataAddress low:  0x40
	// [4] Quantity high:    0x00
	// [5] Quantity low:     0x02
	expected := []byte{0x03, 0x10, 0x00, 0x40, 0x00, 0x02}

	frame, err := mwr.Build()
	if err != nil {
		t.Fatalf("MultipleWritingResponse Build() returned error: %v", err)
	}
	if !reflect.DeepEqual(frame, expected) {
		t.Errorf("MultipleWritingResponse Build() failed.\nExpected: %v\nGot:      %v", expected, frame)
	}
}

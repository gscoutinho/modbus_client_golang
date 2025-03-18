package modbus

import (
	"reflect"
	"testing"
)

func TestTCPRequestWrapper_Build(t *testing.T) {
	// Create a known Modbus frame.
	// For example, a Modbus frame with three bytes: 0x11, 0x22, 0x33.
	modbusFrame := []byte{0x11, 0x22, 0x33}

	// Initialize the TCPRequestWrapper.
	// Note: TransactionID and ProtocolID are not used in the Build() method.
	wrap := &TCPRequestWrapper{
		TransactionID: 1,
		ProtocolID:    0,
		ModbusFrame:   modbusFrame,
	}

	// Build the complete TCP frame.
	frame, err := wrap.Build()
	if err != nil {
		t.Fatalf("Build() returned error: %v", err)
	}

	// Expected frame breakdown:
	// Bytes 0-1: Transaction ID = 0x0001  -> {0x00, 0x01}
	// Bytes 2-3: Protocol ID  = 0x0000  -> {0x00, 0x00}
	// Bytes 4-5: Message Length = length of modbusFrame = 3, in big endian -> {0x00, 0x03}
	// Bytes 6+:  ModbusFrame -> {0x11, 0x22, 0x33}
	expected := []byte{
		0x00, 0x01, // Transaction ID
		0x00, 0x00, // Protocol ID
		0x00, 0x03, // Message Length (3 bytes)
		0x11, 0x22, 0x33, // ModbusFrame payload
	}

	if !reflect.DeepEqual(frame, expected) {
		t.Errorf("Expected frame %v, got %v", expected, frame)
	}

	// Also check that MessageLength was correctly set.
	if wrap.MessageLength != uint16(len(modbusFrame)) {
		t.Errorf("Expected MessageLength %d, got %d", len(modbusFrame), wrap.MessageLength)
	}
}

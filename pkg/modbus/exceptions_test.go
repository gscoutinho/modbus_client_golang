package modbus

import (
	"testing"
)

func TestModbusException_Error(t *testing.T) {
	tests := []struct {
		code     byte
		expected string
	}{
		{
			code:     0x01,
			expected: "Modbus Exception (0x01): Illegal Function: The function code received in the query is not allowable for the slave.",
		},
		{
			code:     0x02,
			expected: "Modbus Exception (0x02): Illegal Data Address: The data address received in the query is not an allowable address for the slave.",
		},
		{
			code:     0x03,
			expected: "Modbus Exception (0x03): Illegal Data Value: A value contained in the query data field is not an allowable value for the slave.",
		},
		{
			code:     0x04,
			expected: "Modbus Exception (0x04): Slave Device Failure: An unrecoverable error occurred while the slave was attempting to perform the requested action.",
		},
		{
			code:     0x05,
			expected: "Modbus Exception (0x05): Acknowledge: The slave has accepted the request and is processing it, but a long duration is required.",
		},
		{
			code:     0x06,
			expected: "Modbus Exception (0x06): Slave Device Busy: The slave is busy processing a long-duration program command.",
		},
		{
			code:     0x07,
			expected: "Modbus Exception (0x07): Negative Acknowledge: The slave cannot perform the requested program function.",
		},
		{
			code:     0x08,
			expected: "Modbus Exception (0x08): Memory Parity Error: The slave detected a parity error in its memory.",
		},
		{
			code:     0x0A,
			expected: "Modbus Exception (0x0A): Gateway Path Unavailable: The gateway was unable to allocate an internal communication path.",
		},
		{
			code:     0x0B,
			expected: "Modbus Exception (0x0B): Gateway Target Device Failed to Respond: No response was obtained from the target device.",
		},
		{
			code:     0xFF,
			expected: "Modbus Exception (0xFF): Unknown Modbus exception",
		},
	}

	for _, tc := range tests {
		err := NewModbusException(tc.code)
		if err == nil {
			t.Errorf("Expected error for code 0x%02X, but got nil", tc.code)
			continue
		}
		if err.Error() != tc.expected {
			t.Errorf("For code 0x%02X, expected error string:\n%s\nGot:\n%s",
				tc.code, tc.expected, err.Error())
		}
	}
}

package modbus

import "fmt"

type ModbusExceptionCode byte

const (
	ExceptionIllegalFunction              ModbusExceptionCode = 0x01
	ExceptionIllegalDataAddress           ModbusExceptionCode = 0x02
	ExceptionIllegalDataValue             ModbusExceptionCode = 0x03
	ExceptionSlaveDeviceFailure           ModbusExceptionCode = 0x04
	ExceptionAcknowledge                  ModbusExceptionCode = 0x05
	ExceptionSlaveDeviceBusy              ModbusExceptionCode = 0x06
	ExceptionNegativeAcknowledge          ModbusExceptionCode = 0x07
	ExceptionMemoryParityError            ModbusExceptionCode = 0x08
	ExceptionGatewayPathUnavailable       ModbusExceptionCode = 0x0A
	ExceptionGatewayTargetFailedToRespond ModbusExceptionCode = 0x0B
)

var exceptionDescriptions = map[ModbusExceptionCode]string{
	ExceptionIllegalFunction:              "Illegal Function: The function code received in the query is not allowable for the slave.",
	ExceptionIllegalDataAddress:           "Illegal Data Address: The data address received in the query is not an allowable address for the slave.",
	ExceptionIllegalDataValue:             "Illegal Data Value: A value contained in the query data field is not an allowable value for the slave.",
	ExceptionSlaveDeviceFailure:           "Slave Device Failure: An unrecoverable error occurred while the slave was attempting to perform the requested action.",
	ExceptionAcknowledge:                  "Acknowledge: The slave has accepted the request and is processing it, but a long duration is required.",
	ExceptionSlaveDeviceBusy:              "Slave Device Busy: The slave is busy processing a long-duration program command.",
	ExceptionNegativeAcknowledge:          "Negative Acknowledge: The slave cannot perform the requested program function.",
	ExceptionMemoryParityError:            "Memory Parity Error: The slave detected a parity error in its memory.",
	ExceptionGatewayPathUnavailable:       "Gateway Path Unavailable: The gateway was unable to allocate an internal communication path.",
	ExceptionGatewayTargetFailedToRespond: "Gateway Target Device Failed to Respond: No response was obtained from the target device.",
}

type ModbusException struct {
	Code ModbusExceptionCode
}

func (e *ModbusException) Error() string {
	desc, ok := exceptionDescriptions[e.Code]
	if !ok {
		desc = "Unknown Modbus exception"
	}
	return fmt.Sprintf("Modbus Exception (0x%02X): %s", e.Code, desc)
}

func NewModbusException(code byte) error {
	return &ModbusException{Code: ModbusExceptionCode(code)}
}

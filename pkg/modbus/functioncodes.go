package modbus

//FC type
type FunctionCode byte

const (
	FCReadCoils               FunctionCode = 1
	FCReadInputStatus         FunctionCode = 2
	FCReadHoldingRegisters    FunctionCode = 3
	FCReadInputRegisters      FunctionCode = 4
	FCForceSingleCoil         FunctionCode = 5
	FCPresetSingleRegister    FunctionCode = 6
	FCForceMultipleCoils      FunctionCode = 15
	FCPresetMultipleRegisters FunctionCode = 16
)

func (fc FunctionCode) String() string {
	switch fc {
	case FCReadCoils:
		return "Read Coils"
	case FCReadInputStatus:
		return "Read Input Status"
	case FCReadHoldingRegisters:
		return "Read Holding Registers"
	case FCReadInputRegisters:
		return "Read Input Registers"
	case FCForceSingleCoil:
		return "Force Single Coil"
	case FCPresetSingleRegister:
		return "Preset Single Register"
	case FCForceMultipleCoils:
		return "Force Multiple Coils"
	case FCPresetMultipleRegisters:
		return "Preset Multiple Registers"
	default:
		return "Unknown Function Code"
	}
}

package modbus

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type WordByteOrder struct {
	ByteOrder binary.ByteOrder
	SwapWords bool
}

func swapWords32(b []byte) []byte {
	if len(b) != 4 {
		return b
	}
	return []byte{b[2], b[3], b[0], b[1]}
}

func swapWords64(b []byte) []byte {
	if len(b) != 8 {
		return b
	}
	// Original order: [w0, w1, w2, w3] where each word is 2 bytes.
	// New order: [w3, w2, w1, w0]
	return []byte{
		b[6], b[7],
		b[4], b[5],
		b[2], b[3],
		b[0], b[1],
	}
}

func Int16ToBytes(value int16, order binary.ByteOrder) []byte {
	b := make([]byte, 2)
	order.PutUint16(b, uint16(value))
	return b
}

func Uint16ToBytes(value uint16, order binary.ByteOrder) []byte {
	b := make([]byte, 2)
	order.PutUint16(b, value)
	return b
}

func HexStringToBytes(hexStr string) ([]byte, error) {
	// Remove any "0x" or "0X" prefix.
	if len(hexStr) >= 2 && (hexStr[:2] == "0x" || hexStr[:2] == "0X") {
		hexStr = hexStr[2:]
	}
	return hex.DecodeString(hexStr)
}

func BinaryStringToBytes(binStr string) ([]byte, error) {
	// Parse the binary string into an unsigned integer.
	n, err := strconv.ParseUint(binStr, 2, 64)
	if err != nil {
		return nil, fmt.Errorf("error parsing binary string: %v", err)
	}
	// Determine how many bytes are needed.
	byteLen := (len(binStr) + 7) / 8
	b := make([]byte, byteLen)
	for i := byteLen - 1; i >= 0; i-- {
		b[i] = byte(n & 0xFF)
		n >>= 8
	}
	return b, nil
}

func Int32ToBytes(value int32, order WordByteOrder) []byte {
	b := make([]byte, 4)
	order.ByteOrder.PutUint32(b, uint32(value))
	if order.SwapWords {
		b = swapWords32(b)
	}
	return b
}

func Uint32ToBytes(value uint32, order WordByteOrder) []byte {
	b := make([]byte, 4)
	order.ByteOrder.PutUint32(b, value)
	if order.SwapWords {
		b = swapWords32(b)
	}
	return b
}

func Int64ToBytes(value int64, order WordByteOrder) []byte {
	b := make([]byte, 8)
	order.ByteOrder.PutUint64(b, uint64(value))
	if order.SwapWords {
		b = swapWords64(b)
	}
	return b
}

func Uint64ToBytes(value uint64, order WordByteOrder) []byte {
	b := make([]byte, 8)
	order.ByteOrder.PutUint64(b, value)
	if order.SwapWords {
		b = swapWords64(b)
	}
	return b
}

func Float32ToBytes(value float32, order WordByteOrder) []byte {
	bits := math.Float32bits(value)
	return Uint32ToBytes(bits, order)
}

func Float64ToBytes(value float64, order WordByteOrder) []byte {
	bits := math.Float64bits(value)
	return Uint64ToBytes(bits, order)
}

func BytesToInt16(data []byte, order binary.ByteOrder) (int16, error) {
	if len(data) != 2 {
		return 0, fmt.Errorf("data length is not 2 bytes: got %d", len(data))
	}
	u16 := order.Uint16(data)
	return int16(u16), nil
}

func BytesToUint16(data []byte, order binary.ByteOrder) (uint16, error) {
	if len(data) != 2 {
		return 0, fmt.Errorf("data length is not 2 bytes: got %d", len(data))
	}
	return order.Uint16(data), nil
}

func BytesToHexString(data []byte) string {
	return hex.EncodeToString(data)
}

func BytesToBinaryString(data []byte, sep bool) string {
	var parts []string
	for _, b := range data {
		parts = append(parts, fmt.Sprintf("%08b", b))
	}
	if sep {
		return strings.Join(parts, " ")
	}
	return strings.Join(parts, "")
}

func BytesToInt32(data []byte, order WordByteOrder) (int32, error) {
	if len(data) != 4 {
		return 0, fmt.Errorf("data length is not 4 bytes: got %d", len(data))
	}
	b := data
	if order.SwapWords {
		b = swapWords32(b)
	}
	u32 := order.ByteOrder.Uint32(b)
	return int32(u32), nil
}

func BytesToUint32(data []byte, order WordByteOrder) (uint32, error) {
	if len(data) != 4 {
		return 0, fmt.Errorf("data length is not 4 bytes: got %d", len(data))
	}
	b := data
	if order.SwapWords {
		b = swapWords32(b)
	}
	return order.ByteOrder.Uint32(b), nil
}

func BytesToInt64(data []byte, order WordByteOrder) (int64, error) {
	if len(data) != 8 {
		return 0, fmt.Errorf("data length is not 8 bytes: got %d", len(data))
	}
	b := data
	if order.SwapWords {
		b = swapWords64(b)
	}
	u64 := order.ByteOrder.Uint64(b)
	return int64(u64), nil
}

func BytesToUint64(data []byte, order WordByteOrder) (uint64, error) {
	if len(data) != 8 {
		return 0, fmt.Errorf("data length is not 8 bytes: got %d", len(data))
	}
	b := data
	if order.SwapWords {
		b = swapWords64(b)
	}
	return order.ByteOrder.Uint64(b), nil
}

func BytesToFloat32(data []byte, order WordByteOrder) (float32, error) {
	u32, err := BytesToUint32(data, order)
	if err != nil {
		return 0, err
	}
	return math.Float32frombits(u32), nil
}

func BytesToFloat64(data []byte, order WordByteOrder) (float64, error) {
	u64, err := BytesToUint64(data, order)
	if err != nil {
		return 0, err
	}
	return math.Float64frombits(u64), nil
}

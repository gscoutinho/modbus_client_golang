package modbus

import (
	"encoding/binary"
	"math"
	"strings"
	"testing"
)

// --------------------
// 16-bit Converters
// --------------------

func TestInt16ToBytesAndBack(t *testing.T) {
	order := binary.BigEndian
	values := []int16{-32768, -1, 0, 1, 32767}
	for _, v := range values {
		b := Int16ToBytes(v, order)
		v2, err := BytesToInt16(b, order)
		if err != nil {
			t.Fatalf("Int16 conversion error for value %d: %v", v, err)
		}
		if v != v2 {
			t.Errorf("Int16 mismatch: expected %d, got %d", v, v2)
		}
	}
}

func TestUint16ToBytesAndBack(t *testing.T) {
	order := binary.BigEndian
	values := []uint16{0, 1, 12345, 65535}
	for _, v := range values {
		b := Uint16ToBytes(v, order)
		v2, err := BytesToUint16(b, order)
		if err != nil {
			t.Fatalf("Uint16 conversion error for value %d: %v", v, err)
		}
		if v != v2 {
			t.Errorf("Uint16 mismatch: expected %d, got %d", v, v2)
		}
	}
}

// --------------------
// String Converters
// --------------------

func TestHexStringConversion(t *testing.T) {
	hexStr := "0xDEADBEEF"
	b, err := HexStringToBytes(hexStr)
	if err != nil {
		t.Fatalf("HexStringToBytes error: %v", err)
	}
	// The expected result should be "deadbeef" (lowercase) without the "0x" prefix.
	expected := "deadbeef"
	result := BytesToHexString(b)
	if result != expected {
		t.Errorf("Hex conversion mismatch: expected %s, got %s", expected, result)
	}
}

func TestBinaryStringConversion(t *testing.T) {
	binStr := "11001010"
	b, err := BinaryStringToBytes(binStr)
	if err != nil {
		t.Fatalf("BinaryStringToBytes error: %v", err)
	}
	// Convert back to binary string without spaces.
	result := BytesToBinaryString(b, false)
	// For a single byte, result should have 8 characters.
	// Trim any leading zeros if necessary.
	result = strings.TrimLeft(result, "0")
	expected := strings.TrimLeft(binStr, "0")
	if result != expected {
		t.Errorf("Binary conversion mismatch: expected %s, got %s", binStr, result)
	}
}

// --------------------
// 32-bit Converters
// --------------------

func TestInt32ToBytesAndBack(t *testing.T) {
	orderNoSwap := WordByteOrder{ByteOrder: binary.BigEndian, SwapWords: false}
	orderSwap := WordByteOrder{ByteOrder: binary.BigEndian, SwapWords: true}
	values := []int32{-2147483648, -1, 0, 1, 2147483647}
	for _, v := range values {
		// Test without word swap.
		b := Int32ToBytes(v, orderNoSwap)
		v2, err := BytesToInt32(b, orderNoSwap)
		if err != nil {
			t.Fatalf("Int32 conversion error (no swap) for %d: %v", v, err)
		}
		if v != v2 {
			t.Errorf("Int32 no swap mismatch: expected %d, got %d", v, v2)
		}

		// Test with word swap.
		bSwap := Int32ToBytes(v, orderSwap)
		v2Swap, err := BytesToInt32(bSwap, orderSwap)
		if err != nil {
			t.Fatalf("Int32 conversion error (swap) for %d: %v", v, err)
		}
		if v != v2Swap {
			t.Errorf("Int32 swap mismatch: expected %d, got %d", v, v2Swap)
		}
	}
}

func TestUint32ToBytesAndBack(t *testing.T) {
	orderNoSwap := WordByteOrder{ByteOrder: binary.BigEndian, SwapWords: false}
	orderSwap := WordByteOrder{ByteOrder: binary.BigEndian, SwapWords: true}
	values := []uint32{0, 1, 1234567890, 4294967295}
	for _, v := range values {
		// Test without word swap.
		b := Uint32ToBytes(v, orderNoSwap)
		v2, err := BytesToUint32(b, orderNoSwap)
		if err != nil {
			t.Fatalf("Uint32 conversion error (no swap) for %d: %v", v, err)
		}
		if v != v2 {
			t.Errorf("Uint32 no swap mismatch: expected %d, got %d", v, v2)
		}

		// Test with word swap.
		bSwap := Uint32ToBytes(v, orderSwap)
		v2Swap, err := BytesToUint32(bSwap, orderSwap)
		if err != nil {
			t.Fatalf("Uint32 conversion error (swap) for %d: %v", v, err)
		}
		if v != v2Swap {
			t.Errorf("Uint32 swap mismatch: expected %d, got %d", v, v2Swap)
		}
	}
}

// --------------------
// 64-bit Converters
// --------------------

func TestInt64ToBytesAndBack(t *testing.T) {
	orderNoSwap := WordByteOrder{ByteOrder: binary.BigEndian, SwapWords: false}
	orderSwap := WordByteOrder{ByteOrder: binary.BigEndian, SwapWords: true}
	values := []int64{-9223372036854775808, -1, 0, 1, 9223372036854775807}
	for _, v := range values {
		// Without swap.
		b := Int64ToBytes(v, orderNoSwap)
		v2, err := BytesToInt64(b, orderNoSwap)
		if err != nil {
			t.Fatalf("Int64 conversion error (no swap) for %d: %v", v, err)
		}
		if v != v2 {
			t.Errorf("Int64 no swap mismatch: expected %d, got %d", v, v2)
		}

		// With swap.
		bSwap := Int64ToBytes(v, orderSwap)
		v2Swap, err := BytesToInt64(bSwap, orderSwap)
		if err != nil {
			t.Fatalf("Int64 conversion error (swap) for %d: %v", v, err)
		}
		if v != v2Swap {
			t.Errorf("Int64 swap mismatch: expected %d, got %d", v, v2Swap)
		}
	}
}

func TestUint64ToBytesAndBack(t *testing.T) {
	orderNoSwap := WordByteOrder{ByteOrder: binary.BigEndian, SwapWords: false}
	orderSwap := WordByteOrder{ByteOrder: binary.BigEndian, SwapWords: true}
	values := []uint64{0, 1, 12345678901234567890, 18446744073709551615}
	for _, v := range values {
		// Without swap.
		b := Uint64ToBytes(v, orderNoSwap)
		v2, err := BytesToUint64(b, orderNoSwap)
		if err != nil {
			t.Fatalf("Uint64 conversion error (no swap) for %d: %v", v, err)
		}
		if v != v2 {
			t.Errorf("Uint64 no swap mismatch: expected %d, got %d", v, v2)
		}

		// With swap.
		bSwap := Uint64ToBytes(v, orderSwap)
		v2Swap, err := BytesToUint64(bSwap, orderSwap)
		if err != nil {
			t.Fatalf("Uint64 conversion error (swap) for %d: %v", v, err)
		}
		if v != v2Swap {
			t.Errorf("Uint64 swap mismatch: expected %d, got %d", v, v2Swap)
		}
	}
}

// --------------------
// Floating Point Converters
// --------------------

func TestFloat32ToBytesAndBack(t *testing.T) {
	orderNoSwap := WordByteOrder{ByteOrder: binary.BigEndian, SwapWords: false}
	orderSwap := WordByteOrder{ByteOrder: binary.BigEndian, SwapWords: true}
	values := []float32{-3.14, 0, 3.14, math.MaxFloat32, math.SmallestNonzeroFloat32}
	for _, v := range values {
		// Without swap.
		b := Float32ToBytes(v, orderNoSwap)
		v2, err := BytesToFloat32(b, orderNoSwap)
		if err != nil {
			t.Fatalf("Float32 conversion error (no swap) for %f: %v", v, err)
		}
		// Compare using a tolerance since floating point may have rounding differences.
		if math.Abs(float64(v-v2)) > 1e-6 {
			t.Errorf("Float32 no swap mismatch: expected %f, got %f", v, v2)
		}

		// With swap.
		bSwap := Float32ToBytes(v, orderSwap)
		v2Swap, err := BytesToFloat32(bSwap, orderSwap)
		if err != nil {
			t.Fatalf("Float32 conversion error (swap) for %f: %v", v, err)
		}
		if math.Abs(float64(v-v2Swap)) > 1e-6 {
			t.Errorf("Float32 swap mismatch: expected %f, got %f", v, v2Swap)
		}
	}
}

func TestFloat64ToBytesAndBack(t *testing.T) {
	orderNoSwap := WordByteOrder{ByteOrder: binary.BigEndian, SwapWords: false}
	orderSwap := WordByteOrder{ByteOrder: binary.BigEndian, SwapWords: true}
	values := []float64{-3.1415926535, 0, 3.1415926535, math.MaxFloat64, math.SmallestNonzeroFloat64}
	for _, v := range values {
		// Without swap.
		b := Float64ToBytes(v, orderNoSwap)
		v2, err := BytesToFloat64(b, orderNoSwap)
		if err != nil {
			t.Fatalf("Float64 conversion error (no swap) for %f: %v", v, err)
		}
		if math.Abs(v-v2) > 1e-9 {
			t.Errorf("Float64 no swap mismatch: expected %f, got %f", v, v2)
		}

		// With swap.
		bSwap := Float64ToBytes(v, orderSwap)
		v2Swap, err := BytesToFloat64(bSwap, orderSwap)
		if err != nil {
			t.Fatalf("Float64 conversion error (swap) for %f: %v", v, err)
		}
		if math.Abs(v-v2Swap) > 1e-9 {
			t.Errorf("Float64 swap mismatch: expected %f, got %f", v, v2Swap)
		}
	}
}

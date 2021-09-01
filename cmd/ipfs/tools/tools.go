package tools

import (
	"encoding/binary"
	"log"
	"math"
)

// Check checks for the error, logs it and sends back the bool value.
func Check(err error, info ...string) bool {
	if err != nil {
		log.Println(info, err.Error())
		return false
	}
	return true
}

// NumToBytes - converts integer to bytes, supports int32 and int64
func NumToBytes(num interface{}) []byte {
	var len int
	switch num.(type) {
	case int64:
		len = 8
		b := make([]byte, len)
		binary.BigEndian.PutUint64(b, uint64(num.(int64)))
		return b
	case float64:
		bits := math.Float64bits(num.(float64))
		b := make([]byte, 8)
		binary.BigEndian.PutUint64(b, bits)
		return b
	case int:
		len = 4
		b := make([]byte, len)
		binary.BigEndian.PutUint32(b, uint32(num.(int)))
		return b
	case int32:
		len = 4
		b := make([]byte, len)
		binary.BigEndian.PutUint32(b, uint32(num.(int32)))
		return b
	}
	return []byte{}
}

// BytesToInt - converts byte array to int
func BytesToInt(arr []byte) int {
	return int(binary.BigEndian.Uint32(arr))
}

// BytesToInt64 - converts byte array to int
func BytesToInt64(arr []byte) int64 {
	return int64(binary.BigEndian.Uint64(arr))
}

// BytesToInt32 - converts byte array to int
func BytesToInt32(arr []byte) int32 {
	return int32(binary.BigEndian.Uint32(arr))
}

// BytesToFloat64 - converts bytes to float64
func BytesToFloat64(bytes []byte) float64 {
	bits := binary.BigEndian.Uint64(bytes)
	float := math.Float64frombits(bits)
	return float
}

// BoolToByte - converts bool to byte
func BoolToByte(b bool) byte {
	if b {
		return 1
	}
	return 0
}

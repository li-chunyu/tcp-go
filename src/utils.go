package tcpip

import (
	"encoding/binary"
)

func writeUint16ToByteBigEndian(v uint16) []byte {
	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, v)
	return b
}

func readUint16FromByteBigEndian(b []byte) uint16 {
	return binary.BigEndian.Uint16(b)
}

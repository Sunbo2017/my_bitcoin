package core

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
)

func Int2Hex(v int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, v)
	if err != nil {
		panic(err)
	}
	return buff.Bytes()
}

func Data2Hash(data []byte) []byte {
	hash := sha256.Sum256(data)
	return hash[:]
}

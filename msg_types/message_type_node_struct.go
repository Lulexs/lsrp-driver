package msg_types

import (
	"bytes"
	"encoding/binary"
)

type MessageTypeNode struct {
	NextPossibleMessages []Message
}

func isAuthType(firstByte byte, messageBytes []byte, expectedFirstByte byte, expectedSpecifier int32) bool {
	return firstByte == expectedFirstByte && binary.BigEndian.Uint32(messageBytes[0:4]) == uint32(expectedSpecifier)
}

func writeParams(writer *bytes.Buffer, keyValueMap map[string]string) {
	for key, value := range keyValueMap {
		writer.WriteString(key)
		writer.WriteByte(0)
		writer.WriteString(value)
		writer.WriteByte(0)
	}
}

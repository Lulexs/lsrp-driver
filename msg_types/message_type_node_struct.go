package msg_types

import (
	"bufio"
	"bytes"
	"encoding/binary"
)

type MessageTypeNode struct {
	NextPossibleMessages []Message
}

func isAuthType(reader *bufio.Reader, firstByte byte, expectedSpecifier int32) (bool, error) {
	fb, err := reader.ReadByte()
	if err != nil || fb != firstByte {
		return false, err
	}

	var length int32
	if err := binary.Read(reader, binary.BigEndian, &length); err != nil {
		return false, err
	}

	var specifier int32
	if err := binary.Read(reader, binary.BigEndian, &specifier); err != nil {
		return false, err
	}

	return specifier == expectedSpecifier, nil
}

func writeParams(writer *bytes.Buffer, keyValueMap map[string]string) {
	for key, value := range keyValueMap {
		writer.WriteString(key)
		writer.WriteByte(0)
		writer.WriteString(value)
		writer.WriteByte(0)
	}
}

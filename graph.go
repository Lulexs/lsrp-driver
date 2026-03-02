package main

import (
	"bufio"
	"encoding/binary"
)

type Message interface {
	IsResponseMessageOfMessageType(reader *bufio.Reader) (bool, error)
	GetFirstByte() byte
	GetNextPossibleMessages() []Message // TODO: REPLACE THIS WITH CTOR FUNCTION
}

type MessageTypeNode struct {
	MessageType          string
	FirstByte            byte
	NextPossibleMessages []MessageTypeNode
}

type StartUpMessage struct {
	MessageTypeNode
}

type ErrorResponse struct {
	MessageTypeNode
}

type AuthenticationOk struct {
	MessageTypeNode
}

type AuthenticationClearTextPassword struct {
	MessageTypeNode
}

func (msg StartUpMessage) GetFirstByte() byte {
	return 0
}

func (msg StartUpMessage) IsResponseOfMessageType(reader *bufio.Reader) (bool, error) {
	return false, nil
}

func (msg StartUpMessage) GetNextPossibleMessages() []Message {

}

func (msg ErrorResponse) GetFirstByte() byte {
	return 'E'
}

func (msg ErrorResponse) IsResponseMessageOfMessageType(reader *bufio.Reader) (bool, error) {
	firstByte, err := reader.ReadByte()
	if err != nil {
		return false, err
	} else if firstByte != msg.GetFirstByte() {
		return false, nil
	}
	return true, nil
}

func (msg AuthenticationOk) GetFirstByte() byte {
	return 'R'
}

func (msg AuthenticationOk) IsResponseMessageOfMessageType(reader *bufio.Reader) (bool, error) {
	return isAuthType(reader, msg.GetFirstByte(), 0)
}

func (msg AuthenticationClearTextPassword) GetFirstByte() byte {
	return 'R'
}

func (msg AuthenticationClearTextPassword) IsResponseMessageOfMessageType(reader *bufio.Reader) (bool, error) {
	return isAuthType(reader, msg.GetFirstByte(), 3)
}

func isAuthType(reader *bufio.Reader, firstByte byte, expectedSpecifier int32) (bool, error) {
	fb, err := reader.ReadByte()
	if err != nil || fb != 'R' {
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

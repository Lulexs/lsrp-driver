package msg_types

import "fmt"

type ErrorResponse struct {
	MessageTypeNode
}

func (msg ErrorResponse) GetDisplayName() string {
	return "ErrorResponse"
}

func (msg ErrorResponse) GetFirstByte() byte {
	return 'E'
}

func (msg ErrorResponse) IsResponseMessageOfMessageType(firstByte byte, msgBytes []byte) bool {
	return firstByte == msg.GetFirstByte()
}

func (msg ErrorResponse) GetNextPossibleMessages() []Message {
	return []Message{}
}

func (msg ErrorResponse) PrintError(msgBytes []byte) {
	prevZero := -1
	for i, x := range msgBytes {
		if x == 0 && prevZero == i-1 {
			return
		}
		if x == 0 && string(msgBytes[prevZero+1]) == "M" {
			fmt.Printf("Error: %v\n", string(msgBytes[prevZero+2:i]))
			return
		} else if x == 0 {
			prevZero = i
		}
	}
}

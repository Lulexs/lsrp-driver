package msg_types

import "bufio"

type ErrorResponse struct {
	MessageTypeNode
}

func (msg ErrorResponse) GetDisplayName() string {
	return "ErrorResponse"
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

func (msg ErrorResponse) GetNextPossibleMessages() []Message {
	return []Message{}
}

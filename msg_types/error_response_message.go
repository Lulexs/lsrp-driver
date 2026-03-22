package msg_types

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

package msg_types

type AuthenticationOk struct {
	MessageTypeNode
}

func (msg AuthenticationOk) GetDisplayName() string {
	return "AuthenticationOk"
}

func (msg AuthenticationOk) GetFirstByte() byte {
	return 'R'
}

func (msg AuthenticationOk) IsResponseMessageOfMessageType(firstByte byte, msgBytes []byte) bool {
	return isAuthType(firstByte, msgBytes, msg.GetFirstByte(), 0)
}

func (msg AuthenticationOk) GetNextPossibleMessages() []Message {
	return []Message{}
}

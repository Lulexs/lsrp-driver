package msg_types

type AuthenticationClearTextPassword struct {
	MessageTypeNode
}

func (msg AuthenticationClearTextPassword) GetDisplayName() string {
	return "AuthenticationClearTextPassword"
}

func (msg AuthenticationClearTextPassword) GetFirstByte() byte {
	return 'R'
}

func (msg AuthenticationClearTextPassword) IsResponseMessageOfMessageType(firstByte byte, msgBytes []byte) bool {
	return isAuthType(firstByte, msgBytes, msg.GetFirstByte(), 3)
}

func (msg AuthenticationClearTextPassword) GetNextPossibleMessages() []Message {
	return []Message{}
}

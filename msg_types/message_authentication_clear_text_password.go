package msg_types

import "bufio"

type AuthenticationClearTextPassword struct {
	MessageTypeNode
}

func (msg AuthenticationClearTextPassword) GetDisplayName() string {
	return "AuthenticationClearTextPassword"
}

func (msg AuthenticationClearTextPassword) GetFirstByte() byte {
	return 'R'
}

func (msg AuthenticationClearTextPassword) IsResponseMessageOfMessageType(reader *bufio.Reader) (bool, error) {
	return isAuthType(reader, msg.GetFirstByte(), 3)
}

func (msg AuthenticationClearTextPassword) GetNextPossibleMessages() []Message {
	return []Message{}
}

package msg_types

import "bufio"

type AuthenticationOk struct {
	MessageTypeNode
}

func (msg AuthenticationOk) GetDisplayName() string {
	return "AuthenticationOk"
}

func (msg AuthenticationOk) GetFirstByte() byte {
	return 'R'
}

func (msg AuthenticationOk) IsResponseMessageOfMessageType(reader *bufio.Reader) (bool, error) {
	return isAuthType(reader, msg.GetFirstByte(), 0)
}

func (msg AuthenticationOk) GetNextPossibleMessages() []Message {
	return []Message{}
}

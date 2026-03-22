package msg_types

type Message interface {
	IsResponseMessageOfMessageType(firstByte byte, msgBytes []byte) bool
	GetFirstByte() byte
	GetNextPossibleMessages() []Message
	GetDisplayName() string
}

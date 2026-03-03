package msg_types

type Message interface {
	IsResponseMessageOfMessageType(firstByte byte) (bool, error)
	GetFirstByte() byte
	GetNextPossibleMessages() []Message
	GetDisplayName() string
}

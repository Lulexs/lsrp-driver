package msg_types

type Message interface {
	IsResponseMessageOfMessageType(firstByte byte, msgBytes []byte) bool
	GetFirstByte() byte
	GetNextPossibleMessages() []Message
	GetDisplayName() string
}

type OutgoingMessage interface {
	BuildMessageContent() []byte
	Message
}

type MessageData struct {
	MapData  map[string]string
	ListData []string
}

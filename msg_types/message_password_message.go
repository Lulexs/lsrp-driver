package msg_types

import (
	"bytes"
	"encoding/binary"
	"sync"
)

type PasswordMessage struct {
	MessageTypeNode
	once sync.Once
}

func (msg *PasswordMessage) GetDisplayName() string {
	return "PasswordMessage"
}

func (msg *PasswordMessage) GetFirstByte() byte {
	return 'p'
}

func (msg *PasswordMessage) IsResponseMessageOfMessageType(firstByte byte, msgBytes []byte) bool {
	return firstByte == msg.GetFirstByte()
}

func (msg *PasswordMessage) GetNextPossibleMessages() []Message {
	msg.once.Do(func() {
		if len(msg.NextPossibleMessages) == 0 {
			msg.NextPossibleMessages = []Message{
				&ErrorResponse{},
				&AuthenticationOk{},
			}
		}
	})

	return msg.NextPossibleMessages
}

func (msg *PasswordMessage) BuildMessageContent() []byte {
	return msg.buildMessageContent(MessageData{
		ListData: []string{"postgres"},
	})
}

func (msg *PasswordMessage) buildMessageContent(data MessageData) []byte {
	buf := new(bytes.Buffer)
	buf.Write(make([]byte, 5))

	writeListParams(buf, data.ListData)

	byteArray := buf.Bytes()
	byteArray[0] = msg.GetFirstByte()
	binary.BigEndian.PutUint32(byteArray[1:5], uint32(len(byteArray)-1))

	return byteArray
}

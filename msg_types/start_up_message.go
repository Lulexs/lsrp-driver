package msg_types

import (
	"bytes"
	"encoding/binary"
	"sync"
)

// First message sent by frontend to backend

type StartUpMessage struct {
	MessageTypeNode
	once sync.Once
}

func (msg *StartUpMessage) IsResponseMessageOfMessageType(firstByte byte, msgBytes []byte) bool {
	return false
}

func (msg *StartUpMessage) GetFirstByte() byte {
	return 0
}

func (msg *StartUpMessage) GetNextPossibleMessages() []Message {
	msg.once.Do(func() {
		if len(msg.NextPossibleMessages) == 0 {
			msg.NextPossibleMessages = []Message{
				&ErrorResponse{},
				&AuthenticationClearTextPassword{},
				&AuthenticationOk{},
			}
		}
	})

	return msg.NextPossibleMessages
}

func (msg *StartUpMessage) GetDisplayName() string {
	return "StartUpMessage"
}

func (msg *StartUpMessage) BuildMessageContent() []byte {
	return msg.buildMessageContent(MessageData{MapData: map[string]string{
		"user":     "postgres",
		"database": "postgres",
	}})
}

func (msg *StartUpMessage) buildMessageContent(data MessageData) []byte {
	var majorVer uint16 = 3
	var minorVer uint16 = 2

	buf := new(bytes.Buffer)
	buf.Write(make([]byte, 8))

	writeMapParams(buf, data.MapData)
	buf.WriteByte(0)

	byteArray := buf.Bytes()
	binary.BigEndian.PutUint32(byteArray[0:4], uint32(len(byteArray)))
	binary.BigEndian.PutUint16(byteArray[4:6], majorVer)
	binary.BigEndian.PutUint16(byteArray[6:8], minorVer)

	return byteArray
}

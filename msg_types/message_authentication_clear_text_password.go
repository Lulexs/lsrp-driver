package msg_types

import "sync"

type AuthenticationClearTextPassword struct {
	MessageTypeNode
	once sync.Once
}

func (msg *AuthenticationClearTextPassword) GetDisplayName() string {
	return "AuthenticationClearTextPassword"
}

func (msg *AuthenticationClearTextPassword) GetFirstByte() byte {
	return 'R'
}

func (msg *AuthenticationClearTextPassword) IsResponseMessageOfMessageType(firstByte byte, msgBytes []byte) bool {
	return isAuthType(firstByte, msgBytes, msg.GetFirstByte(), 3)
}

func (msg *AuthenticationClearTextPassword) GetNextPossibleMessages() []Message {
	msg.once.Do(func() {
		if len(msg.NextPossibleMessages) == 0 {
			msg.NextPossibleMessages = []Message{
				&PasswordMessage{},
			}
		}
	})
	return msg.NextPossibleMessages
}

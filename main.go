package main

import (
	"encoding/binary"
	"fmt"
	"lsrp-driver/msg_types"
	"net"
)

func sendMessage(conn net.Conn, byteArray []byte, message msg_types.Message) error {
	fmt.Printf("Sending %v with content %v\n", message.GetDisplayName(), byteArray)

	n, err := conn.Write(byteArray)
	if err != nil || n != len(byteArray) {
		fmt.Printf("Failed to send %v\n", message.GetDisplayName())
		return err
	}
	fmt.Printf("Successfully sent %v\n", message.GetDisplayName())
	return nil
}

func receiveMessage(conn net.Conn) (byte, []byte, error) {
	recvBuffer := make([]byte, 5)
	_, err := conn.Read(recvBuffer)
	if err != nil {
		return 0, nil, err
	}
	msgLen := binary.BigEndian.Uint32(recvBuffer[1:5])
	restBuffer := make([]byte, msgLen-4)
	n, err := conn.Read(restBuffer)
	if err != nil || msgLen-4 != uint32(n) {
		return 0, nil, err
	}

	return recvBuffer[0], restBuffer, nil
}

func main() {

	conn, err := net.Dial("tcp", "localhost:5432")

	if err != nil {
		panic(err)
	}

	var outgoingMsg msg_types.Message = &msg_types.StartUpMessage{}
	content := outgoingMsg.(msg_types.OutgoingMessage).BuildMessageContent()

	for {
		err = sendMessage(conn, content, outgoingMsg)
		if err != nil {
			conn.Close()
		}

		firstByte, restBuffer, err := receiveMessage(conn)
		if err != nil {
			fmt.Println(err)
			return
		}

		found := false
		for _, msg := range outgoingMsg.GetNextPossibleMessages() {
			if msg.IsResponseMessageOfMessageType(firstByte, restBuffer) {
				fmt.Printf("Received %v with content %v\n", msg.GetDisplayName(), restBuffer)
				if errResponse, ok := msg.(*msg_types.ErrorResponse); ok {
					errResponse.PrintError(restBuffer)
					conn.Close()
					return
				}
				
				possibleResponses := msg.GetNextPossibleMessages()
				if len(possibleResponses) != 1 {
					panic("Found 0 or more than 1 possible response to message")
				} else if _, ok := possibleResponses[0].(msg_types.OutgoingMessage); !ok {
					panic("Found impossible outgoing message type")
				}
				outgoingMsg = possibleResponses[0]
				content = possibleResponses[0].(msg_types.OutgoingMessage).BuildMessageContent()

				found = true
				break
			}
		}

		if !found {
			panic("Unexpected message arrived from server")

		}
	}

}

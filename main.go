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
	restBuffer := make([]byte, msgLen)
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

	startupMsg := &msg_types.StartUpMessage{}
	content := startupMsg.BuildMessageContent(map[string]string{
		"username": "postgres",
		"database": "postgres",
	})

	err = sendMessage(conn, content, startupMsg)
	if err != nil {
		conn.Close()
	}

	firstByte, restBuffer, err := receiveMessage(conn)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, msg := range startupMsg.GetNextPossibleMessages() {
		if msg.IsResponseMessageOfMessageType(firstByte, restBuffer) {
			fmt.Println(msg.GetDisplayName())
		}
	}

}

package main

import (
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

	recvBuffer := make([]byte, 1)
	n, err := conn.Read(recvBuffer)
	if err != nil {
		fmt.Println("Failed to receive any response")
		return
	}

	fmt.Printf("Received %v bytes\n", n)
	fmt.Println(recvBuffer[:n])

}

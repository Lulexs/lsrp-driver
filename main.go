package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
)

func main() {

	conn, err := net.Dial("tcp", "localhost:5432")

	if err != nil {
		panic(err)
	}

	startupMessage := buildStartupMessage()
	fmt.Println("Startup message: ", startupMessage)
	n, err := conn.Write(startupMessage)

	if err != nil {
		fmt.Println("Failed to send startup message")
		return
	}
	fmt.Printf("Successfully sent %v bytes\n", n)

	recvBuffer := make([]byte, 1024)
	n, err = conn.Read(recvBuffer)
	if err != nil {
		fmt.Println("Failed to receive response to startup message")
		return
	}
	fmt.Printf("Received %v bytes\n", n)
	fmt.Println(recvBuffer[:n])

}

func writeParams(writer *bytes.Buffer, keyValueMap map[string]string) {
	for key, value := range keyValueMap {
		writer.WriteString(key)
		writer.WriteByte(0)
		writer.WriteString(value)
		writer.WriteByte(0)
	}
}

func buildStartupMessage() []byte {

	params := map[string]string{
		"username": "postgres",
		"database": "postgres",
	}

	var majorVer uint16 = 3
	var minorVer uint16 = 2

	byteArray := make([]byte, 8)
	buf := bytes.NewBuffer(byteArray)

	binary.BigEndian.PutUint32(byteArray[0:4], 0)
	binary.BigEndian.PutUint16(byteArray[4:6], majorVer)
	binary.BigEndian.PutUint16(byteArray[6:8], minorVer)

	writeParams(buf, params)
	buf.WriteByte(0)

	binary.BigEndian.PutUint32(byteArray[0:4], uint32(len(byteArray)))

	return byteArray

}

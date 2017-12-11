package server

import (
	"os"
	"net"
	"bufio"
	"fmt"
)

func addSizeToMessage(bytes []byte)([]byte){
	size := intTo2bytes(len(bytes))

	bytes = append(size, bytes...)

	return bytes
}

func intTo2bytes(len int)[]byte{
	size := make([]byte, 2)
	size[0] = (byte)(len % 256)
	size[1] = (byte)(len / 256)

	return size
}

func main(){
	args := os.Args[1:]

	name := args[0]
	server := args[1]

	conn, _ := net.Dial("tcp", server)

	writer := bufio.NewWriter(conn)
	reader := bufio.NewReader(conn)

	writer.Write(([]byte)(name))
	writer.Flush()

	fmt.Println("Written name")

	msgBytes := make([]byte, 2)
	msgBytes[0] = 0
	msgBytes[1] = 0
	msgBytes = append(msgBytes, ([]byte)("00000001")...)
	msgBytes = append(msgBytes, 0)
	msgBytes = append(msgBytes, 0)
	msgBytes = append(msgBytes, ([]byte)("00000000")...)
	msgBytes = append(msgBytes, 0)
	msgBytes = append(msgBytes, 0)
	msgBytes = append(msgBytes, ([]byte)("00000002")...)
	msgBytes = append(msgBytes, 0)
	msgBytes = append(msgBytes, 0)
	msgBytes = append(msgBytes, ([]byte)("00000001")...)
	msgBytes = append(msgBytes, 0)
	msgBytes = append(msgBytes, 0)
	msgBytes = append(msgBytes, ([]byte)("00000000")...)
	msgBytes = append(msgBytes, 0)
	msgBytes = append(msgBytes, 0)
	msgBytes = append(msgBytes, ([]byte)("10000000")...)
	msgBytes = append(msgBytes, ([]byte)("ukryta wiadomosc")...)

	msgBytes = addSizeToMessage(msgBytes)

	fmt.Println(string(msgBytes))
	writer.Write(msgBytes)
	writer.Flush()
	fmt.Println(writer.Buffered())

	fmt.Println("Written msg")

	msgOk := make([]byte, 100)
	reader.Read(msgOk)

	fmt.Println("MSGOK")

	reader.Read(msgOk)
	fmt.Println(string(msgOk))

	conn.Close()
	conn.Close()
}
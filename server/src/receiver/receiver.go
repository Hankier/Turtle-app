package receiver

import (
	"net"
	"sync"
	"bufio"
	"io"
	"log"
	"messageHandler"
)

type Receiver struct{
	socket net.Conn
	messageHandler messageHandler.MessageHandler
}

func NewReceiver(socket net.Conn, messageHandler messageHandler.MessageHandler)(*Receiver){
	recv := new(Receiver)
	recv.socket = socket
	recv.messageHandler = messageHandler

	return recv
}

func (recv *Receiver)Loop(wg *sync.WaitGroup){
	defer wg.Done()

	log.Print("Starting receiver loop")


	reader := bufio.NewReader(recv.socket)

	size := make([]byte, 2)
	for {
		_, err := io.ReadFull(recv.socket, size)
		if err != nil{log.Print("Receiver " + err.Error());break}

		n := twoBytesToInt(size)

		log.Print("Receiving ", n, " bytes")

		bytes := make([]byte, n)
		_, err = io.ReadFull(reader, bytes)
		if err != nil{log.Print("Receiver " + err.Error());break}

		log.Print("Received ", n, " bytes:", string(bytes))

		recv.messageHandler.HandleBytes(bytes)
	}
}

func twoBytesToInt(size []byte)int{
	num := 0

	num += (int)(size[0])
	num += (int)(size[1]) * 256

	return num
}

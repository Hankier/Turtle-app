package receiver

import (
	"net"
	"sync"
	"bufio"
	"io"
	"log"
	"messageHandler"
	"utils"
)

type Receiver struct{
	sessionName string
	socket net.Conn
	messageHandler messageHandler.MessageHandler
}

func NewReceiver(sessionName string, socket net.Conn, messageHandler messageHandler.MessageHandler)(*Receiver){
	recv := new(Receiver)
	recv.sessionName = sessionName
	recv.socket = socket
	recv.messageHandler = messageHandler

	return recv
}

func (recv *Receiver)Loop(wg *sync.WaitGroup){
	defer wg.Done()

	reader := bufio.NewReader(recv.socket)

	size := make([]byte, 2)
	for {
		_, err := io.ReadFull(recv.socket, size)
		if err != nil{log.Print("Receiver " + err.Error());break}

		n := utils.TwoBytesToInt(size)

		bytes := make([]byte, n)
		_, err = io.ReadFull(reader, bytes)
		if err != nil{log.Print("Receiver " + err.Error());break}

		recv.messageHandler.HandleBytes(recv.sessionName, bytes)
	}
}

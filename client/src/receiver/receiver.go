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
	socket      net.Conn
	msghandler  messageHandler.MessageHandler
}

func New(name string, socket net.Conn, msghandler messageHandler.MessageHandler)(*Receiver){
	recv := new(Receiver)
	recv.sessionName = name
	recv.socket = socket
	recv.msghandler = msghandler

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

		//log.Print("DEBUG Received from: " + recv.sessionName)
		//log.Print("DEBUG Received msg: " + (string)(bytes))

		recv.msghandler.HandleBytes(recv.sessionName, bytes)
	}
}

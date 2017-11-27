package session

import (
	"net"
	"sync"
	"log"
)

type Session struct{
	name string
	socket net.Conn
	sender *sender.Sender
	receiver *receiver.Receiver
	wg sync.WaitGroup
}

func NewSession(socket net.Conn, name string, messageHandler messageHandler.MessageHandler)(*Session){
	session := new(Session)

	session.name = name
	session.socket = socket
	session.sender = sender.NewSender(socket)
	session.receiver = receiver.NewReceiver(socket, messageHandler)
	session.wg.Add(2)

	return session
}

func (session *Session)Start(){
	log.Print("Starting session " + session.name)

	go sender.Loop(session.wg)
	go receiver.Loop(session.wg)

	session.wg.Wait()

	log.Print("Session ended " + session.name)
}

func (session *Session)DeleteSession(){
	session.socket.Close()
}

func (session *Session)Send(bytes []byte){
	session.sender.send(bytes)
}

func (session *Session)UnlockSending(){
	session.sender.UnlockSending()
}

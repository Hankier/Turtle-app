package session

import (
	"net"
	"sync"
	"log"
	"sender"
	"receiver"
	"messageHandler"
)

type Session struct{
	name string
	socket *net.Conn
	sender *sender.SenderImpl
	receiver *receiver.Receiver
	wg *sync.WaitGroup
}

func NewSession(socket *net.Conn, name string, messageHandler messageHandler.MessageHandler)(*Session){
	session := new(Session)

	session.name = name
	session.socket = socket
	session.sender = sender.NewSenderImpl(socket)
	session.receiver = receiver.NewReceiver(socket, messageHandler)
	session.wg = &sync.WaitGroup{}
	session.wg.Add(1)

	return session
}

func (session *Session)Start(){
	log.Print("Starting session " + session.name)

	go session.sender.Loop(session.wg)
	go session.receiver.Loop(session.wg)

	session.wg.Wait()

	session.DeleteSession()

	log.Print("Session ended " + session.name)
}

func (session *Session)DeleteSession(){
	session.sender.Stop()
	(*session.socket).Close()
}

func (session *Session)Send(bytes []byte){
	bytes = append(([]byte)(session.name), bytes...)
	session.sender.Send(bytes)
}

func (session *Session)UnlockSending(){
	session.sender.UnlockSending()
}

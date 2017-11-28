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
	wgS *sync.WaitGroup
	wgR *sync.WaitGroup
}

func NewSession(socket *net.Conn, name string, messageHandler messageHandler.MessageHandler)(*Session){
	session := new(Session)

	session.name = name
	session.socket = socket
	session.sender = sender.NewSenderImpl(socket)
	session.receiver = receiver.NewReceiver(socket, messageHandler)
	session.wgS = &sync.WaitGroup{}
	session.wgS.Add(1)
	session.wgR = &sync.WaitGroup{}
	session.wgR.Add(1)

	return session
}

func (session *Session)Start(){
	log.Print("Starting session " + session.name)

	go session.sender.Loop(session.wgS)
	go session.receiver.Loop(session.wgR)

	session.wgR.Wait()
	session.DeleteSession()
	session.wgS.Wait()

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

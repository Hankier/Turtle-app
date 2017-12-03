package session

import (
	"net"
	"sync"
	"log"
	"sender"
	"receiver"
	"messageHandler"
	"sessionHandler"
	"message"
)

type Session struct{
	name string
	socket net.Conn
	sender *sender.SenderImpl
	receiver *receiver.Receiver
	handler sessionHandler.SessionHandler
	wgS *sync.WaitGroup
	wgR *sync.WaitGroup
}

func NewSession(socket net.Conn, name string, messageHandler messageHandler.MessageHandler, handler sessionHandler.SessionHandler)(*Session){
	session := new(Session)

	session.name = name
	session.socket = socket
	session.sender = sender.NewSenderImpl(socket)
	session.receiver = receiver.NewReceiver(name, socket, messageHandler)
	session.handler = handler
	session.wgS = &sync.WaitGroup{}
	session.wgS.Add(1)
	session.wgR = &sync.WaitGroup{}
	session.wgR.Add(1)

	return session
}

func (session *Session)Start(){
	defer session.socket.Close()
	log.Print("Starting session: " + session.name)

	go session.sender.Loop(session.wgS)
	go session.receiver.Loop(session.wgR)

	session.wgR.Wait()
	session.sender.Stop()
	session.wgS.Wait()

	session.handler.RemoveSession()

	log.Print("Session ended: " + session.name)
}

func (session *Session)DeleteSession(){
	session.socket.Close()
}

func (session *Session)Send(msg *message.Message){
	log.Print("Sending to: " + session.name)
	session.sender.Send(msg)
}

func (session *Session)SendInstant(msg *message.Message){
	log.Print("Sending instant to: " + session.name)
	session.sender.SendInstant(msg)
}

func (session *Session)UnlockSending(){
	log.Print("Unlock sending to: " + session.name)
	session.sender.UnlockSending()
}

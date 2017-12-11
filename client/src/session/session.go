package session

import (
	"net"
	"sync"
	"log"
	"sender"
	"receiver"
	"messageHandler"
	"message"
	"session/handler"
)

type Session struct{
	name     string
	socket   net.Conn
	sender   *sender.SenderImpl
	recver   *receiver.Receiver
	handler  handler.Handler
	wgSender *sync.WaitGroup
	wgRecver *sync.WaitGroup
}

func New(socket net.Conn, name string, messageHandler messageHandler.MessageHandler, handler handler.Handler)(*Session){
	s := new(Session)

	s.name = name
	s.socket = socket
	s.sender = sender.New(socket)
	s.recver = receiver.New(name, socket, messageHandler)
	s.handler = handler
	s.wgSender = &sync.WaitGroup{}
	s.wgSender.Add(1)
	s.wgRecver = &sync.WaitGroup{}
	s.wgRecver.Add(1)

	return s
}

func (session *Session)Start(){
	defer session.socket.Close()
	log.Print("Starting session: " + session.name)

	go session.sender.Loop(session.wgSender)
	go session.recver.Loop(session.wgRecver)

	session.wgRecver.Wait()
	session.sender.Stop()
	session.wgSender.Wait()

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

func (session *Session)GetName()string{
	return session.name
}

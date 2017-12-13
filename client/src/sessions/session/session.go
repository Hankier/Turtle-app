package session

import (
	"net"
	"sync"
	"log"
	"sessions"
)

type Session struct{
	name     string
	socket   net.Conn

	wgSender *sync.WaitGroup
	msgsSent [][]byte
	msgsMutex sync.Mutex
	canSend  bool
	canSendMutex sync.Mutex
	stopped bool

	wgRecver *sync.WaitGroup
	sessionsRecver sessions.Receiver

	handler  sessions.Handler
}

func New(socket net.Conn, name string, sessionsRecver sessions.Receiver, handler sessions.Handler)(*Session){
	s := new(Session)

	s.name = name
	s.socket = socket

	s.wgSender = &sync.WaitGroup{}
	s.wgSender.Add(1)
	s.msgsSent = make([][]byte, 0, 10)
	s.msgsMutex = sync.Mutex{}
	s.canSend = true
	s.canSendMutex = sync.Mutex{}
	s.stopped = false

	s.wgRecver = &sync.WaitGroup{}
	s.wgRecver.Add(1)
	s.sessionsRecver = sessionsRecver

	s.handler = handler

	return s
}

func (s *Session)Start(){
	defer s.socket.Close()
	log.Print("Starting session: " + s.name)

	go s.SendLoop(s.wgSender)
	go s.ReceiveLoop(s.wgRecver)

	s.wgRecver.Wait()
	s.sender.Stop()
	s.wgSender.Wait()

	s.handler.RemoveSession()

	log.Print("Session ended: " + s.name)
}

func (s *Session)DeleteSession(){
	s.socket.Close()
}

func (s *Session)GetName()string{
	return s.name
}

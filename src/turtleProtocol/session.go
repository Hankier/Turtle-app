package turtleProtocol

import (
	"net"
	"sync"
	"log"
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

	wgReceiver       *sync.WaitGroup

	recv Receiver
}

func NewSession(socket net.Conn, name string, recv Receiver)(*Session){
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

	s.wgReceiver = &sync.WaitGroup{}
	s.wgReceiver.Add(1)
	s.recv = recv

	return s
}

func (s *Session)Start(){
	defer s.socket.Close()
	log.Print("Starting session: " + s.name)

	go s.SendLoop()
	go s.ReceiveLoop()

	s.wgReceiver.Wait()
	s.Stop()
	s.wgSender.Wait()

	log.Print("Session ended: " + s.name)
}

func (s *Session)DeleteSession(){
	s.socket.Close()
}

func (s *Session)GetName()string{
	return s.name
}

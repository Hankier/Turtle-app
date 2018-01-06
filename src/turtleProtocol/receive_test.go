package turtleProtocol

import (
	"net"
	"time"
	"errors"
	"testing"
	"sync"
)

type SocketMockR struct{
	readstate int
	readconst byte
}

func (s *SocketMockR)Read(b []byte) (n int, err error){
	switch s.readstate{
	case 0:
		b[0] = 1
		b[1] = 0
		s.readstate++
		return 2, nil
	case 1:
		b[0] = s.readconst
		s.readstate++
		return 1, nil
	case 2:
		return 0, errors.New("EOF")
	}
	return 0, nil
}

func (s *SocketMockR)Write(b []byte) (n int, err error){
	return 0, nil
}

func (s *SocketMockR)Close() error{
	return nil
}

func (s *SocketMockR)LocalAddr() net.Addr{
	return nil
}

func (s *SocketMockR)RemoteAddr() net.Addr{
	return nil
}

func (s *SocketMockR)SetDeadline(t time.Time) error{
	return nil
}

func (s *SocketMockR)SetReadDeadline(t time.Time) error{
	return nil
}

func (s *SocketMockR)SetWriteDeadline(t time.Time) error{
	return nil
}

type ReceiverMock struct{
	from	string
	handled []byte
}

func (h *ReceiverMock)OnReceive(from string, bytes []byte){
	h.from = from
	h.handled = bytes
}

func TestReceiver(t *testing.T) {
	sessionName := "testsess"
	socket := &SocketMockR{0, 14}
	reciever := &ReceiverMock{}

	wg := &sync.WaitGroup{}
	wg.Add(1)

	session := New(socket, sessionName, reciever, nil)
	if session.sessionsReceiver != reciever ||
		session.name != sessionName ||
		session.socket != socket{

		t.Error("Constructor error")
	}

	session.ReceiveLoop()

	if reciever.handled[0] != socket.readconst ||
		reciever.from != sessionName{

		t.Error("Bad handled data")
	}
}

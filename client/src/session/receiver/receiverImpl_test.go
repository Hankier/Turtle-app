package receiver

import (
	"testing"
	"net"
	"time"
	"sync"
	"errors"
)

type SocketMock struct{
	readstate int
	readconst byte
}

func (s *SocketMock)Read(b []byte) (n int, err error){
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

func (s *SocketMock)Write(b []byte) (n int, err error){
	return 0, nil
}

func (s *SocketMock)Close() error{
	return nil
}

func (s *SocketMock)LocalAddr() net.Addr{
	return nil
}

func (s *SocketMock)RemoteAddr() net.Addr{
	return nil
}

func (s *SocketMock)SetDeadline(t time.Time) error{
	return nil
}

func (s *SocketMock)SetReadDeadline(t time.Time) error{
	return nil
}

func (s *SocketMock)SetWriteDeadline(t time.Time) error{
	return nil
}

type MessageHandlerMock struct{
	from	string
	handled []byte
}

func (h *MessageHandlerMock)HandleBytes(from string, bytes []byte){
	h.from = from
	h.handled = bytes
}

func TestReceiver(t *testing.T) {
	sessionName := "testsess"
	socket := &SocketMock{0, 14}
	msghandler := &MessageHandlerMock{}
	wg := &sync.WaitGroup{}
	wg.Add(1)

	recv := New(sessionName, socket, msghandler)
	if recv.msghandler != msghandler ||
		recv.sessionName != sessionName ||
		recv.socket != socket{

		t.Error("Constructor error")
	}

	recv.Loop(wg)
	if msghandler.handled[0] != socket.readconst ||
		msghandler.from != sessionName{

		t.Error("Bad handled data")
	}
}

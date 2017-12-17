package session

import (
	"net"
	"time"
	"errors"
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

type MessageHandlerMock struct{
	from	string
	handled []byte
}

func (h *MessageHandlerMock)HandleBytes(from string, bytes []byte){
	h.from = from
	h.handled = bytes
}

/*func TestReceiver(t *testing.T) {
	sessionName := "testsess"
	socket := &SocketMockR{0, 14}
	msghandler := &MessageHandlerMock{}
	wg := &sync.WaitGroup{}
	wg.Add(1)

	recv := New(socket, sessionName, nil, nil)
*//*	if recv.msghandler != msghandler ||
		recv.sessionName != sessionName ||
		recv.socket != socket{

		t.Error("Constructor error")
	}*//*

	recv.ReceiveLoop()
	if msghandler.handled[0] != socket.readconst ||
		msghandler.from != sessionName{

		t.Error("Bad handled data")
	}
}*///TODO

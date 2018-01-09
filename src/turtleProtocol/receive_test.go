package turtleProtocol

import (
	"net"
	"time"
	"errors"
	"testing"
	"sync"
	"turtleProtocol/msg"
	"crypt"
)

type SocketMockR struct{
	readstate int
	readconst byte
	packet []byte
}

func (s *SocketMockR)Read(b []byte) (n int, err error){
	switch s.readstate{
	case 0:
		b[0] = s.packet[0]
		b[1] = s.packet[1]
		b[2] = s.packet[2]
		b[3] = s.packet[3]
		s.readstate++
		return 4, nil
	case 1:
		b[0] = s.packet[4]
		b[1] = s.packet[5]
		s.readstate++
		return 2, nil
	case 2:
		b[0] = s.packet[6]
		b[1] = s.packet[7]
		b[2] = s.packet[8]
		s.readstate++
		return 3, nil
	case 3:
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
	message *msg.Message
}

func (h *ReceiverMock)OnReceive(message *msg.Message){
	h.message = message
}

func TestReceiver(t *testing.T) {
	sessionName := "testsess"
	socketConst := 14
	message := new(msg.Message)
	message.SetEncType(crypt.PLAIN)
	message.SetMessageType(msg.DEFAULT)
	message.SetMessageContent([]byte{(byte)(socketConst)})
	packet := addSizeToPacket(addSizeToMessage(message.ToBytes()))

	socket := &SocketMockR{0, (byte)(socketConst), packet}

	recv := &ReceiverMock{}

	wg := &sync.WaitGroup{}
	wg.Add(1)

	session := NewSession(socket, sessionName, recv)
	if session.name != sessionName || session.socket != socket{
		t.Error("Constructor error")
	}

	session.ReceiveLoop()

	if recv.message.GetMessageContent()[0] != socket.readconst{

		t.Error("Bad handled data")
	}
}

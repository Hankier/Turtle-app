package sender

import (
	"testing"
	"message"
	"net"
	"time"
	"bytes"
	_"log"
)

type SocketMock struct{
	writedata []byte
}

func (s *SocketMock)Read(b []byte) (n int, err error){
	return 0, nil
}

func (s *SocketMock)Write(b []byte) (n int, err error){
	s.writedata = b
	return len(b), nil
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

func TestSenderImpl_messagesToSingle(t *testing.T) {
	bytes := make([][]byte, 2)
	bytes[0] = make([]byte, 1)
	bytes[0][0] = 15
	bytes[1] = make([]byte, 1)
	bytes[1][0] = 14

	result := messagesToSingle(bytes)

	if result[0] != 15 || result[1] != 14{
		t.Error("Expected 15 14, got ", result[0], " ", result[1])
	}
}

func TestSenderImpl_Send(t *testing.T) {
	socket := &SocketMock{}
	sender := New(socket)

	msg := message.NewMessageOK()
	sender.Send(msg)

	hasmsg := false
	msgbytes := addSizeToBytes(msg.ToBytes())

	for i := 0; i < len(sender.msgs); i++{
		if bytes.Compare(sender.msgs[i], msgbytes) == 0{
			hasmsg = true
			break
		}
	}

	if !hasmsg{
		t.Error("sender doesnt have test message")
	}
}

func TestSenderImpl_SendInstant(t *testing.T) {
	socket := &SocketMock{}
	sender := New(socket)

	msg := message.NewMessageOK()
	sender.SendInstant(msg)

	msgbytes := addSizeToBytes(msg.ToBytes())

	if bytes.Compare(socket.writedata, msgbytes) != 0{
		t.Error("test message not sent")
	}
}

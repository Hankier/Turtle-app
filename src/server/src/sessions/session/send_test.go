package session

import (
	"testing"
	"msgs/msg"
	"net"
	"time"
	"bytes"
	_"log"
)

type SendSocketMock struct{
	writedata []byte
}

func (s *SendSocketMock)Read(b []byte) (n int, err error){
	return 0, nil
}

func (s *SendSocketMock)Write(b []byte) (n int, err error){
	s.writedata = b
	return len(b), nil
}

func (s *SendSocketMock)Close() error{
	return nil
}

func (s *SendSocketMock)LocalAddr() net.Addr{
	return nil
}

func (s *SendSocketMock)RemoteAddr() net.Addr{
	return nil
}

func (s *SendSocketMock)SetDeadline(t time.Time) error{
	return nil
}

func (s *SendSocketMock)SetReadDeadline(t time.Time) error{
	return nil
}

func (s *SendSocketMock)SetWriteDeadline(t time.Time) error{
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
	socket := &SendSocketMock{}
	sender := New(socket)

	msg := msg.NewMessageOK()
	sender.Send(msg.ToBytes())

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
	socket := &SendSocketMock{}
	sender := New(socket)

	msg := msg.NewMessageOK()
	sender.SendInstant(msg.ToBytes())

	msgbytes := addSizeToBytes(msg.ToBytes())

	if bytes.Compare(socket.writedata, msgbytes) != 0{
		t.Error("test message not sent")
	}
}

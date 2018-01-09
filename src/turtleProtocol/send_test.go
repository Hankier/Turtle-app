package turtleProtocol

import (
	"testing"
	"net"
	"time"
	"bytes"
	_"log"
	"turtleProtocol/msg"
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
	content := make([][]byte, 2)
	content[0] = make([]byte, 1)
	content[0][0] = 15
	content[1] = make([]byte, 1)
	content[1][0] = 14

	result := messagesToSingle(content)

	if result[0] != 15 || result[1] != 14{
		t.Error("Expected 15 14, got ", result[0], " ", result[1])
	}
}

func TestSenderImpl_Send(t *testing.T) {
	socket := &SocketMock{}
	sender := NewSession(socket, "", nil)


	message, _ := msg.FromBytes([]byte{1, 0, 20, 123})
	sender.Send(message)

	hasmsg := false
	msgbytes := addSizeToMessage(message.ToBytes())

	for i := 0; i < len(sender.msgsSent); i++{
		if bytes.Compare(sender.msgsSent[i], msgbytes) == 0{
			hasmsg = true
			break
		}
	}

	if !hasmsg{
		t.Error("sender doesnt have test message")
	}
}

func TestSenderImpl_Send_Multiple_Msgs(t *testing.T) {
	socket := &SocketMock{}
	sender := NewSession(socket, "", nil)


	message, _ := msg.FromBytes([]byte{1, 0, 20, 123})
	message2, _ := msg.FromBytes([]byte{1, 0, 20, 123, 50})
	sender.Send(message)
	sender.Send(message2)

	hasmsg := 0
	msgbytes := addSizeToMessage(message.ToBytes())
	msgbytes2 := addSizeToMessage(message2.ToBytes())

	for i := 0; i < len(sender.msgsSent); i++{
		if bytes.Compare(sender.msgsSent[i], msgbytes) == 0{
			hasmsg++
		}
		if bytes.Compare(sender.msgsSent[i], msgbytes2) == 0{
			hasmsg++
		}
	}
	if hasmsg != 2{
		t.Error("sender doesnt have test message")
	}
}


func TestSenderImpl_Send_Multiple_Msgs2(t *testing.T) {
	socket := &SocketMock{}
	sender := NewSession(socket, "", nil)


	message, _ := msg.FromBytes([]byte{1, 0, 20, 123})
	message2, _ := msg.FromBytes([]byte{1, 0, 20, 123, 50})
	sender.Send(message)
	sender.Send(message2)

	msgbytes := addSizeToMessage(message.ToBytes())
	msgbytes2 := addSizeToMessage(message2.ToBytes())

	finalMessage := messagesToSingle(sender.msgsSent)

	var expectedMessage []byte
	expectedMessage = append(expectedMessage, msgbytes ...)
	expectedMessage = append(expectedMessage, msgbytes2 ...)

	if bytes.Compare(finalMessage, expectedMessage) != 0{
		t.Error("sender doesnt have test message")
	}
}

func TestSenderImpl_SendInstant(t *testing.T) {
	socket := &SocketMock{}
	sender := NewSession(socket, "",nil)

	sender.SendConfirmation()

	msgbytes := addSizeToPacket(addSizeToMessage(msg.NewMessageOK().ToBytes()))

	if bytes.Compare(socket.writedata, msgbytes) != 0{
		t.Error("test message not sent")
	}
}

func Test_IntToFourBytes(t *testing.T){
	number := 123456789

	expected := []byte{21, 205, 91, 7}

	result := intToFourBytes(number)

	if !bytes.Equal(expected, result){
		t.Fail()
	}
}

package turtleProtocol

import (
	"testing"
	"net"
	"time"
	"bytes"
	_"log"
	"fmt"
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
	sender := NewSession(socket, "", nil)


	msg := []byte{1, 0, 20, 123}
	sender.Send(msg)


	hasmsg := false
	msgbytes := addSizeToBytes(msg)

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


	msg := []byte{1, 0, 20, 123}
	msg2 := []byte{1, 0, 20, 123, 50}
	sender.Send(msg)
	sender.Send(msg2)

	hasmsg := 0
	msgbytes := addSizeToBytes(msg)
	msgbytes2 := addSizeToBytes(msg2)

	for i := 0; i < len(sender.msgsSent); i++{
		if bytes.Compare(sender.msgsSent[i], msgbytes) == 0{
			hasmsg++
			fmt.Println(sender.msgsSent[i])
		}
		if bytes.Compare(sender.msgsSent[i], msgbytes2) == 0{
			hasmsg++
			fmt.Println(sender.msgsSent[i])

		}
	}
	fmt.Println(hasmsg)
	if hasmsg != 2{
		t.Error("sender doesnt have test message")
	}
}


func TestSenderImpl_Send_Multiple_Msgs2(t *testing.T) {
	socket := &SocketMock{}
	sender := NewSession(socket, "", nil)


	msg := []byte{1, 0, 20, 123}
	msg2 := []byte{1, 0, 20, 123, 50}
	sender.Send(msg)
	sender.Send(msg2)

	msgbytes := addSizeToBytes(msg)
	msgbytes2 := addSizeToBytes(msg2)

	finalMessage := messagesToSingle(sender.msgsSent)

	var expectedMessage []byte
	expectedMessage = append(expectedMessage, msgbytes ...)
	expectedMessage = append(expectedMessage, msgbytes2 ...)
	fmt.Println(expectedMessage)
	fmt.Println(finalMessage)
	if bytes.Compare(finalMessage, expectedMessage) != 0{
		t.Error("sender doesnt have test message")
	}
}

func TestSenderImpl_SendInstant(t *testing.T) {
	socket := &SocketMock{}
	sender := NewSession(socket, "",nil)

	msg := []byte{1, 0}
	sender.SendInstant(msg)

	msgbytes := addSizeToBytes(msg)

	if bytes.Compare(socket.writedata, msgbytes) != 0{
		t.Error("test message not sent")
	}
}

func Test_IntToFourBytes(t *testing.T){
	number := 123456789

	expected := []byte{21, 205, 91, 7}

	result := intToFourBytes(number)

	fmt.Println(expected)
	fmt.Println(result)

	if !bytes.Equal(expected, result){
		t.Fail()
	}
}

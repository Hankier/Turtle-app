package session

import (
	"bufio"
	"io"
	"log"
	"reflect"
)

func (s *Session)ReceiveLoop(){
	defer s.wgReceiver.Done()

	reader := bufio.NewReader(s.socket)

	size := make([]byte, 2)
	for {
		_, err := io.ReadFull(s.socket, size)
		if err != nil{log.Print(reflect.TypeOf(s).String() + err.Error());break}

		n := twoBytesToInt(size)

		bytes := make([]byte, n)
		_, err = io.ReadFull(reader, bytes)
		if err != nil{log.Print(reflect.TypeOf(s).String() + err.Error());break}

		//log.Print("DEBUG Received from: " + recv.sessionName)
		//log.Print("DEBUG Received msg: " + (string)(bytes))

		s.sessionsReceiver.OnReceive(s.name, bytes)
	}
}

func twoBytesToInt(size []byte)int{
	num := 0

	num += (int)(size[0])
	num += (int)(size[1]) * 256

	return num
}

package sender

import (
	"net"
	"sync"
	"bufio"
	"time"
	"message"
)

const LOOP_TIME = time.Second

type SenderImpl struct{
	socket       net.Conn
	msgs         [][]byte
	msgsmutex    *sync.Mutex
	cansend      bool
	cansendmutex *sync.Mutex
	stopped      bool
}

func New(socket net.Conn)(*SenderImpl){
	sender := new(SenderImpl)

	sender.socket = socket
	sender.msgs = make([][]byte, 0, 10)
	sender.msgsmutex = &sync.Mutex{}
	sender.cansend = true
	sender.cansendmutex = &sync.Mutex{}
	sender.stopped = false

	return sender
}

func (s *SenderImpl)Loop(wg *sync.WaitGroup){
	defer wg.Done()

	writer := bufio.NewWriter(s.socket)

	for !s.stopped{
		s.cansendmutex.Lock()
		if !s.cansend {
			s.cansendmutex.Unlock()
		} else {
			s.cansendmutex.Unlock()
			s.msgsmutex.Lock()
			if len(s.msgs) > 0{
				s.cansend = false
				messagesCopy := s.msgs[:]
				s.msgs = s.msgs[:0]
				s.msgsmutex.Unlock()

				packet := messagesToSingle(messagesCopy)

				_, err := writer.Write(packet)
				if err != nil{s.stopped = true; break}
				err = writer.Flush()
				if err != nil{s.stopped = true; break}
			} else {
				s.msgsmutex.Unlock()
			}
		}
		time.Sleep(LOOP_TIME)
	}
}

func messagesToSingle(bytes [][]byte)([]byte){
	result := make([]byte, 0)
	for i := 0; i < len(bytes); i++{
		result = append(result, bytes[i]...)
	}
	return result
}

func (s *SenderImpl)Stop(){
	s.stopped = true
}

func (s *SenderImpl)Send(msg *message.Message){

	content := msg.ToBytes()
	content = addSizeToBytes(content)

	s.msgsmutex.Lock()
	s.msgs = append(s.msgs, content)
	s.msgsmutex.Unlock()
}

func (s *SenderImpl)SendInstant(msg *message.Message){
	content := msg.ToBytes()
	content = addSizeToBytes(content)
	s.socket.Write(content)
}

func (s *SenderImpl)UnlockSending(){
	s.cansend = true
}


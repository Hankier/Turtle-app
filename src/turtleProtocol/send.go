package turtleProtocol

import (
	"time"
	"bufio"
	"log"
)

const LOOP_TIME = time.Second

func (s *Session)SendLoop(){
	defer s.wgSender.Done()

	writer := bufio.NewWriter(s.socket)

	for !s.stopped{
		s.canSendMutex.Lock()
		if !s.canSend {
			s.canSendMutex.Unlock()
		} else {
			s.canSendMutex.Unlock()
			s.msgsMutex.Lock()
			if len(s.msgsSent) > 0{
				s.canSend = false
				messagesCopy := s.msgsSent[:]
				s.msgsSent = s.msgsSent[:0]
				s.msgsMutex.Unlock()

				packet := messagesToSingle(messagesCopy)

				_, err := writer.Write(packet)
				if err != nil{s.stopped = true; break}
				err = writer.Flush()
				if err != nil{s.stopped = true; break}
			} else {
				s.msgsMutex.Unlock()
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

func (s *Session)Stop(){
	s.stopped = true
}

func (s *Session)Send(content []byte){
	log.Print("Sending to: " + s.name)

	content = addSizeToBytes(content)

	s.msgsMutex.Lock()
	s.msgsSent = append(s.msgsSent, content)
	s.msgsMutex.Unlock()
}

func (s *Session)SendInstant(content []byte){
	log.Print("Sending instant to: " + s.name)

	content = addSizeToBytes(content)
	s.socket.Write(content)
}

func (s *Session)UnlockSending(){
	log.Print("Unlock sending to: " + s.name)

	s.canSend = true
}

func addSizeToBytes(bytes []byte)([]byte){
	size := intToTwobytes(len(bytes))

	bytes = append(size, bytes...)

	return bytes
}

func intToTwobytes(len int)[]byte{
	size := make([]byte, 2)
	size[0] = (byte)(len % 256)
	size[1] = (byte)(len / 256)

	return size
}



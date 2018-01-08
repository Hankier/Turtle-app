package turtleProtocol

import (
	"time"
	"bufio"
	"log"
	_"fmt"
	"turtleProtocol/msg"
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

				messagesInSingle := messagesToSingle(messagesCopy)

				packet := addSizeToPacket(messagesInSingle)

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

func messagesToSingle(content [][]byte)([]byte){
	log.Print("Compressing ", len(content), " messages")
	result := make([]byte, 0)
	for i := 0; i < len(content); i++{
		result = append(result, content[i]...)
	}
	return result
}

func (s *Session)Stop(){
	s.stopped = true
}

func (s *Session)Send(message *msg.Message){
	log.Print("Sending to: " + s.name)

	content := message.ToBytes()

	content = addSizeToMessage(content)

	s.msgsMutex.Lock()
	s.msgsSent = append(s.msgsSent, content)
	s.msgsMutex.Unlock()
}

func (s *Session) SendConfirmation(){
	log.Print("Sending response to: " + s.name)

	content := msg.NewMessageOK().ToBytes()
	content = addSizeToMessage(content)

	packet := addSizeToPacket(content)

	s.socket.Write(packet)
}

func (s *Session)UnlockSending(){
	log.Print("Unlock sending to: " + s.name)

	s.canSend = true
}

func addSizeToMessage(content []byte)([]byte){
	size := intToTwobytes(len(content))

	content = append(size, content...)

	return content
}

func addSizeToPacket(content []byte)([]byte){
	size := intToFourBytes(len(content))

	content = append(size, content...)

	return content
}

func intToTwobytes(len int)[]byte{
	size := make([]byte, 2)
	size[0] = (byte)(len % 256)
	size[1] = (byte)(len / 256)

	return size
}

func intToFourBytes(num int)[]byte{
	bytes := make([]byte, 4)
	bytes[0] = (byte)(num % 256)
	num /= 256
	bytes[1] = (byte)(num % 256)
	num /= 256
	bytes[2] = (byte)(num % 256)
	num /= 256
	bytes[3] = (byte)(num % 256)

	return bytes
}


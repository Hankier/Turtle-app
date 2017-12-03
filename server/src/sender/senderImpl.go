package sender

import (
	"net"
	"sync"
	"bufio"
	"time"
	"utils"
	"message"
)

const LOOP_TIME = time.Second

type SenderImpl struct{
	socket net.Conn
	messages [][]byte
	messagesMutex *sync.Mutex
	canSend bool
	canSendMutex *sync.Mutex
	stopped bool
}

func NewSenderImpl(socket net.Conn)(*SenderImpl){
	sender := new(SenderImpl)

	sender.socket = socket
	sender.messages = make([][]byte, 0, 10)
	sender.messagesMutex = &sync.Mutex{}
	sender.canSend = true
	sender.canSendMutex = &sync.Mutex{}
	sender.stopped = false

	return sender
}

func (sender *SenderImpl)Loop(wg *sync.WaitGroup){
	defer wg.Done()

	writer := bufio.NewWriter(sender.socket)

	for !sender.stopped{
		sender.canSendMutex.Lock()
		if !sender.canSend{
			sender.canSendMutex.Unlock()
		} else {
			sender.canSendMutex.Unlock()
			sender.messagesMutex.Lock()
			if len(sender.messages) > 0{
				sender.canSend = false
				messagesCopy := sender.messages[:]
				sender.messages = sender.messages[:0]
				sender.messagesMutex.Unlock()

				packet := messagesToSingle(messagesCopy)
				writer.Write(packet)
				writer.Flush()
			} else {
				sender.messagesMutex.Unlock()
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

func (sender *SenderImpl)Stop(){
	sender.stopped = true
}

func (sender *SenderImpl)Send(msg *message.Message){
	bytes := msg.ToBytes()
	bytes = addSizeToMessage(bytes)

	sender.messagesMutex.Lock()
	sender.messages = append(sender.messages, bytes)
	sender.messagesMutex.Unlock()
}

func (sender *SenderImpl)SendInstant(msg *message.Message){
	bytes := msg.ToBytes()
	bytes = addSizeToMessage(bytes)
	sender.socket.Write(bytes)
}

func (sender *SenderImpl)UnlockSending(){
	sender.canSend = true
}

func addSizeToMessage(bytes []byte)([]byte){
	size := utils.IntToTwobytes(len(bytes))

	bytes = append(size, bytes...)

	return bytes
}


package sender

import (
	"net"
	"sync"
	"bufio"
	"time"
	"log"
	"strconv"
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

	log.Print("Starting sender loop")

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
				log.Print("Sender sending message length " + strconv.Itoa(len(packet)) + " - " + string(packet))
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

func (sender *SenderImpl)Send(bytes []byte){
	bytes = addSizeToMessage(bytes)

	sender.messagesMutex.Lock()
	sender.messages = append(sender.messages, bytes)
	sender.messagesMutex.Unlock()
}

func (sender *SenderImpl)SendInstant(bytes []byte){
	bytes = addSizeToMessage(bytes)
	sender.socket.Write(bytes)
}

func (sender *SenderImpl)UnlockSending(){
	sender.canSend = true
}

func addSizeToMessage(bytes []byte)([]byte){
	size := intTo2bytes(len(bytes))

	bytes = append(size, bytes...)

	return bytes
}

func intTo2bytes(len int)[]byte{
	size := make([]byte, 2)
	size[0] = (byte)(len % 256)
	size[1] = (byte)(len / 256)

	return size
}

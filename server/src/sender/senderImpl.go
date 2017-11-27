package sender

import (
	"net"
	"sync"
)

const LOOP_TIME int = 1

type SenderImpl struct{
	socket *net.Conn
	messages [][]byte
	messagesMutex *sync.Mutex
	canSend bool
	stopped bool
}

func NewSenderImpl(socket *net.Conn)(*SenderImpl){
	sender := new(SenderImpl)

	sender.socket = socket
	sender.messages = make([][]byte, 10)
	sender.messagesMutex = &sync.Mutex{}
	sender.canSend = true
	sender.stopped = false

	return sender
}

func (sender *SenderImpl)Loop(wg sync.WaitGroup){
	//TODO
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

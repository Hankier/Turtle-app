package messageHandler

import (
	"sessionsSender"
	"decrypter"
)

type TYPE byte

const (
	MSG TYPE = iota
	MSG_OK
	PING
)

type Message struct{
	messageType TYPE
	previousName string
	messageContent []byte
}

func FromBytes(bytes []byte)(*Message){
	//no previous name and type
	if len(bytes) < 9{
		return nil
	}
	msg := new(Message)

	msg.previousName = string(bytes[0:8])
	msg.messageType = (TYPE)(bytes[8])
	msg.messageContent = append([]byte(nil), bytes[9:]...)

	return msg
}

func (msg *Message)toBytes()[]byte{
	length := len(msg.messageContent) + 1 //+TYPE
	bytes := make([]byte, length)

	bytes[0] = (byte)(msg.messageType)
	for i := 0; i < len(msg.messageContent); i++{
		bytes[i + 1] = msg.messageContent[i]
	}

	return bytes
}

func (msg *Message)Handle(sender sessionsSender.SessionsSender, decrypter decrypter.Decrypter){
	msg.messageContent = decrypter.Decrypt(msg.messageContent)

	switch msg.messageType{
	case MSG:
		msg.handleMSG(sender)
		break
	case MSG_OK:
		msg.handleMSG_OK(sender)
		break
	}
}

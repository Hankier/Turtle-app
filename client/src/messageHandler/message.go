package messageHandler

import (
	"cryptographer"
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
	encType cryptographer.TYPE
	messageContent []byte
}

func FromBytes(from string, bytes []byte)(*Message){
	//no previous name and type
	if len(bytes) < 1{
		return nil
	}
	msg := new(Message)

	msg.previousName = from
	msg.messageType = (TYPE)(bytes[0])
	msg.encType = (cryptographer.TYPE)(bytes[1])
	msg.messageContent = append([]byte(nil), bytes[2:]...)

	return msg
}

func (msg *Message)ToBytes()[]byte{
	length := len(msg.messageContent) + 2 //+TYPE +ENC TYPE
	bytes := make([]byte, length)

	bytes[0] = (byte)(msg.messageType)
	bytes[1] = (byte)(msg.encType)
	for i := 0; i < len(msg.messageContent); i++{
		bytes[i + 2] = msg.messageContent[i]
	}

	return bytes
}

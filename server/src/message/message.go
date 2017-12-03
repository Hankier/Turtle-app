package message

import (
	"cryptographer"
	"utils"
)

type TYPE byte

const (
	MSG TYPE = iota
	MSG_OK
	PING
)

type Message struct{
	messageType    TYPE
	previousName   string
	encType        cryptographer.TYPE
	messageContent []byte
}

func (msg *Message)GetMessageType() TYPE{
	return msg.messageType
}
func (msg *Message)GetPreviousName() string{
	return msg.previousName
}
func (msg *Message)GetEncType() cryptographer.TYPE{
	return msg.encType
}
func (msg *Message)GetMessageContent() []byte{
	return msg.messageContent
}

func (msg *Message)SetMessageType(messageType TYPE){
	msg.messageType = messageType
}
func (msg *Message)SetPreviousName(previousName string){
	msg.previousName = previousName
}
func (msg *Message)SetEncType(encType cryptographer.TYPE){
	msg.encType = encType
}
func (msg *Message)SetMessageContent(messageContent []byte){
	msg.messageContent = messageContent
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
	length := len(msg.messageContent) + 2 //+SIZE +TYPE +ENC TYPE
	bytes := make([]byte, length)

	bytes[0] = (byte)(msg.messageType)
	bytes[1] = (byte)(msg.encType)
	for i := 0; i < len(msg.messageContent); i++{
		bytes[i + 2] = msg.messageContent[i]
	}

	bytes = addSizeToBytes(bytes)

	return bytes
}

func addSizeToBytes(bytes []byte)([]byte){
	size := utils.IntToTwobytes(len(bytes))

	bytes = append(size, bytes...)

	return bytes
}

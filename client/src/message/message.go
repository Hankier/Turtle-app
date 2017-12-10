package message

import (
	"cryptographer"
)

type TYPE byte

const (
	DEFAULT TYPE = iota
	OK
	PING
)

type Message struct{
	messageType    TYPE
	encType        cryptographer.TYPE
	messageContent []byte
}

func NewMessage(msgT TYPE, encT cryptographer.TYPE, msgc []byte)(*Message){
	msg := new(Message)
	msg.messageContent = msgc
	msg.messageType = msgT
	msg.encType = encT
	return msg
}

func (msg *Message)GetMessageType() TYPE{
	return msg.messageType
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
func (msg *Message)SetEncType(encType cryptographer.TYPE){
	msg.encType = encType
}
func (msg *Message)SetMessageContent(messageContent []byte){
	msg.messageContent = messageContent
}

func FromBytes(bytes []byte)(*Message){
	//no previous name and type
	if len(bytes) < 1{
		return nil
	}
	msg := new(Message)

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

	return bytes
}

func NewMessageOK()(*Message){
	msg := new(Message)
	msg.encType = cryptographer.PLAIN;
	msg.messageType = OK;

	return msg;
}

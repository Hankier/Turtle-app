package msg

import (
	"crypt"
	"errors"
)

type TYPE byte

const (
	DEFAULT TYPE = iota
	OK
	PING
)

type Message struct{
	messageType    TYPE
	encType        crypt.TYPE
	messageContent []byte
}

func New(msgT TYPE, encT crypt.TYPE, msgc []byte)(*Message){
	msg := new(Message)
	msg.messageContent = msgc
	msg.messageType = msgT
	msg.encType = encT
	return msg
}

func (msg *Message)GetMessageType() TYPE{
	return msg.messageType
}
func (msg *Message)GetEncType() crypt.TYPE{
	return msg.encType
}
func (msg *Message)GetMessageContent() []byte{
	return msg.messageContent
}

func (msg *Message)SetMessageType(messageType TYPE){
	msg.messageType = messageType
}
func (msg *Message)SetEncType(encType crypt.TYPE){
	msg.encType = encType
}
func (msg *Message)SetMessageContent(messageContent []byte){
	msg.messageContent = messageContent
}

func FromBytes(bytes []byte)(*Message, error){
	//no previous name and type
	if len(bytes) < 1{
		return nil, errors.New( "Message.FromBytes: Too few bytes to construct a message")
	}
	msg := new(Message)

	msg.messageType = (TYPE)(bytes[0])
	msg.encType = (crypt.TYPE)(bytes[1])
	msg.messageContent = append([]byte(nil), bytes[2:]...)

	return msg, nil
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
	msg.encType = crypt.PLAIN;
	msg.messageType = OK;

	return msg;
}

package msg

import (
	"errors"
)

type TYPE byte

const (
	DEFAULT TYPE = iota
	OK
	ADD
	REMOVE
)

type Message struct{
	messageType    TYPE
	messageContent []byte
}

func New(msgT TYPE, msgc []byte)(*Message){
	msg := new(Message)
	msg.messageContent = msgc
	msg.messageType = msgT
	return msg
}

func (msg *Message)GetMessageType() TYPE{
	return msg.messageType
}

func (msg *Message)GetMessageContent() []byte{
	return msg.messageContent
}

func (msg *Message)SetMessageType(messageType TYPE){
	msg.messageType = messageType
}

func (msg *Message)SetMessageContent(messageContent []byte){
	msg.messageContent = messageContent
}

func FromBytes(bytes []byte)(*Message, error){
	//no previous name and type
	if len(bytes) < 1{
		return nil, errors.New("Too few bytes to construct a message")
	}
	msg := new(Message)

	msg.messageType = (TYPE)(bytes[0])
	msg.messageContent = append([]byte(nil), bytes[1:]...)

	return msg, nil
}
func (msg *Message)ToBytes()[]byte{
	length := len(msg.messageContent) + 1 //+SIZE +TYPE +ENC TYPE
	bytes := make([]byte, length)

	bytes[0] = (byte)(msg.messageType)
	for i := 0; i < len(msg.messageContent); i++{
		bytes[i + 1] = msg.messageContent[i]
	}

	return bytes
}

func NewMessageOK()(*Message){
	msg := new(Message)
	msg.messageType = OK;

	return msg;
}

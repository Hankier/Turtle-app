package message

import (
	"cryptographer"
	"utils"
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
	encType        cryptographer.TYPE
	messageContent []byte
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

func FromBytes(bytes []byte)(*Message, error){
	//no previous name and type
	if len(bytes) < 2{
		return nil, errors.New("no message type and/or encryption type")
	}
	msg := new(Message)

	msg.messageType = (TYPE)(bytes[0])
	msg.encType = (cryptographer.TYPE)(bytes[1])
	msg.messageContent = append([]byte(nil), bytes[2:]...)

	return msg, nil
}

func (msg *Message)ToBytes()[]byte{
	bytes := make([]byte, len(msg.messageContent))

	copy(bytes, msg.messageContent)

	bytes = addSizeToBytes(bytes)

	return bytes
}

func addSizeToBytes(bytes []byte)([]byte){
	size := utils.IntToTwobytes(len(bytes))

	bytes = append(size, bytes...)

	return bytes
}

func BuildMessageOK()(*Message){
	msg := new(Message)
	msg.messageContent = make([]byte, 2)
	msg.messageContent[0] = (byte)(OK);
	msg.messageContent[1] = (byte)(cryptographer.PLAIN);

	return msg;
}

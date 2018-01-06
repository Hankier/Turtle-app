package msg

import (
	"errors"
	"client/convos/convo/key"
)

type TYPE byte

const (
	DEFAULT TYPE = iota
	INIT_DATA
	COMMON_KEY_PROTOCOL
)

type ConversationMessage struct{
	msgtype    TYPE
	enctype    key.TYPE
	msgcontent []byte
}

func New(msgtype TYPE, enctype key.TYPE, content []byte)*ConversationMessage{
	convoMsg := new(ConversationMessage)
	convoMsg.msgtype = msgtype
	convoMsg.enctype = enctype
	convoMsg.msgcontent = content
	return convoMsg
}

func FromBytes(content []byte)(*ConversationMessage, error){
	//no previous name and type
	if len(content) < 1{
		return nil, errors.New("Convo.Msg.FromBytes: Too few bytes to create a ConversationMessage")
	}
	msg := new(ConversationMessage)

	msg.msgtype = (TYPE)(content[0])
	msg.enctype = (key.TYPE)(content[1])
	msg.msgcontent = append([]byte(nil), content[2:]...)

	return msg, nil
}

func (msg *ConversationMessage)ToBytes()[]byte{

	length := len(msg.msgcontent) + 2 //+SIZE +TYPE +ENC TYPE
	content := make([]byte, length)

	content[0] = (byte)(msg.msgtype)
	content[1] = (byte)(msg.enctype)
	for i := 0; i < len(msg.msgcontent); i++{
		content[i + 2] = msg.msgcontent[i]
	}

	return content
}

func (msg *ConversationMessage)GetMessageType()TYPE{
	return msg.msgtype
}

func (msg *ConversationMessage)GetEncryptionType()key.TYPE{
	return msg.enctype
}

func (msg *ConversationMessage)GetMessageContent()[]byte{
	return msg.msgcontent
}
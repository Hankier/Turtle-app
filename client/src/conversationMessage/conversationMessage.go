package conversationMessage

import (
	"commonKeyProtocol"
)

type TYPE byte

const (
	DEFAULT TYPE = iota
	INIT_DATA
	COMMON_KEY_PROTOCOL
)

type ConversationMessage struct{
	messageType TYPE
	encType commonKeyProtocol.TYPE
	messageContent []byte
}

func NewConversationMessage(messageType TYPE, encType commonKeyProtocol.TYPE, content []byte)*ConversationMessage{
	convoMsg := new(ConversationMessage)
	convoMsg.messageType = messageType
	convoMsg.encType = encType
	convoMsg.messageContent = content
	return convoMsg
}

func FromBytes(bytes []byte)*ConversationMessage{
	//no previous name and type
	if len(bytes) < 1{
		return nil
	}
	msg := new(ConversationMessage)

	msg.messageType = (TYPE)(bytes[0])
	msg.encType = (commonKeyProtocol.TYPE)(bytes[1])
	msg.messageContent = append([]byte(nil), bytes[2:]...)

	return msg
}

func (msg *ConversationMessage)ToBytes()[]byte{

	length := len(msg.messageContent) + 2 //+SIZE +TYPE +ENC TYPE
	bytes := make([]byte, length)

	bytes[0] = (byte)(msg.messageType)
	bytes[1] = (byte)(msg.encType)
	for i := 0; i < len(msg.messageContent); i++{
		bytes[i + 2] = msg.messageContent[i]
	}

	return bytes
}
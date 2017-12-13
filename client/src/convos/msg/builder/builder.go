package builder

import (
	"convos/msg"
	"commonKeyProtocol"
)

type ConversationMessageBuilder struct{
	commonKeyProtocol commonKeyProtocol.CommonKeyProtocol

	messageType    msg.TYPE
	encType        commonKeyProtocol.TYPE
	messageContent []byte
}

func New(commonKeyProt commonKeyProtocol.CommonKeyProtocol)*ConversationMessageBuilder{
	builder := new(ConversationMessageBuilder)
	builder.commonKeyProtocol = commonKeyProt
	builder.messageType = msg.DEFAULT
	builder.encType = commonKeyProtocol.PLAIN
	builder.messageContent = make([]byte, 0, 0)
	return builder
}

func (builder *ConversationMessageBuilder)SetEncryption(encType commonKeyProtocol.TYPE){
	builder.encType = encType
}

func (builder *ConversationMessageBuilder)SetMessage(message string){
	builder.messageContent = ([]byte)(message)
}

func (builder *ConversationMessageBuilder)SetCommonKeyData(part int, content []byte){
	builder.messageContent = builder.commonKeyProtocol.GetCommonKeyData(part)
}

func (builder *ConversationMessageBuilder)SetInitData(){
	//TODO get this data from nodeCrypto
}

func (builder *ConversationMessageBuilder)Build()[]byte{
	convoMsg := msg.New(builder.messageType, builder.encType, builder.messageContent)
	return convoMsg.ToBytes()
}

func (builder *ConversationMessageBuilder)ParseCommand(message string){
	builder.SetMessage(message)
	//TODO
}

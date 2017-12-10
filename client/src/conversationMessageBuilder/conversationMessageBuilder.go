package conversationMessageBuilder

import (
	"receiverKeyHandler"
	"conversationMessage"
	"commonKeyProtocol"
)

type ConversationMessageBuilder struct{
	commonKeyProtocol commonKeyProtocol.CommonKeyProtocol
	receiverKeyHandler receiverKeyHandler.ReceiverKeyHandler

	messageType conversationMessage.TYPE
	encType commonKeyProtocol.TYPE
	messageContent []byte
}

func NewConversationMessageBuilder(commonKeyProt commonKeyProtocol.CommonKeyProtocol, receiverKeyHandl receiverKeyHandler.ReceiverKeyHandler)*ConversationMessageBuilder{
	builder := new(ConversationMessageBuilder)
	builder.commonKeyProtocol = commonKeyProt
	builder.receiverKeyHandler = receiverKeyHandl
	builder.messageType = conversationMessage.DEFAULT
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
	convoMsg := conversationMessage.NewConversationMessage(builder.messageType, builder.encType, builder.messageContent)
	return convoMsg.ToBytes()
}

func (builder *ConversationMessageBuilder)ParseString(message string){
	//TODO
}

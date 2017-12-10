package conversationMessageBuilder

import (
	"commonKeyProtocol"
	"receiverKeyHandler"
	"conversationMessage"
	"cryptographer"
)

type ConversationMessageBuilder struct{
	commonKeyProtocol commonKeyProtocol.CommonKeyProtocol
	receiverKeyHandler receiverKeyHandler.ReceiverKeyHandler

	messageType conversationMessage.TYPE
	encType commonKeyProtocol.TYPE
	messageContent []byte
}

func NewConversationMessageBuilder(commonKeyProtocol commonKeyProtocol.CommonKeyProtocol, receiverKeyHandler receiverKeyHandler.ReceiverKeyHandler)*ConversationMessageBuilder{
	builder := new(ConversationMessageBuilder)
	builder.commonKeyProtocol = commonKeyProtocol
	builder.receiverKeyHandler = receiverKeyHandler
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
	builder.commonKeyProtocol.SetCommonKeyData(part, content)
}

func (builder *ConversationMessageBuilder)SetInitData(content []byte){
	if len(content) > 0 {
		typ := (cryptographer.TYPE)(content[0])
		keyData := content[1:]
		builder.receiverKeyHandler.SetKey(typ, keyData)
	}
}

func (builder *ConversationMessageBuilder)Build()[]byte{
	convoMsg := conversationMessage.NewConversationMessage(builder.messageType, builder.encType, builder.messageContent)
	return convoMsg.ToBytes()
}

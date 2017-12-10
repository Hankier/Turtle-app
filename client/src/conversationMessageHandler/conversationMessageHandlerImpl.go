package conversationMessageHandler

import (
	"commonKeyProtocol"
	"receiverKeyHandler"
	"textReceiver"
	"conversationMessage"
)

type ConversationMessageHandlerImpl struct{
	commonKeyProt commonKeyProtocol.CommonKeyProtocol
	recvKeyHandler receiverKeyHandler.ReceiverKeyHandler
	textRecv textReceiver.TextReceiver
}

func NewConversationMessageHandlerImpl(commonKeyProt commonKeyProtocol.CommonKeyProtocol, recvKeyHandler receiverKeyHandler.ReceiverKeyHandler, textRecv textReceiver.TextReceiver)(*ConversationMessageHandlerImpl){
	convMHI := new(ConversationMessageHandlerImpl)
	convMHI.commonKeyProt = commonKeyProt
	convMHI.recvKeyHandler = recvKeyHandler
	convMHI.textRecv = textRecv
	return convMHI
}

func (convMHI *ConversationMessageHandlerImpl)HandleBytes(from string, bytes []byte){
	msg := conversationMessage.FromBytes(bytes)

	convMHI.handle(from, msg)
}

func (convMHI *ConversationMessageHandlerImpl)handle(from string, msg *conversationMessage.ConversationMessage){

	decrypted := convMHI.commonKeyProt.Decrypt(msg.GetMessageContent())

	switch(msg.GetMessageType()){
	case conversationMessage.DEFAULT:
		convMHI.handleDEFAULT(from, decrypted)
	case conversationMessage.COMMON_KEY_PROTOCOL:
		convMHI.handleCOMMON_KEY_PROTOCOL(decrypted)
	case conversationMessage.INIT_DATA:
		convMHI.handleINIT_DATA(decrypted)
	}
}
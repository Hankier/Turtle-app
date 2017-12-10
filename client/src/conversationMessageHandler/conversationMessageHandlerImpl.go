package conversationMessageHandler

import (
	"commonKeyProtocol"
	"receiverKeyHandler"
	"textReceiver"
	"conversationMessage"
)

type ConversationMessageHandlerImpl struct{
	commonKeyProtocol commonKeyProtocol.CommonKeyProtocol
	receiverKeyHandler receiverKeyHandler.ReceiverKeyHandler
	textReceiver textReceiver.TextReceiver
}

func (convMHI *ConversationMessageHandlerImpl)HandleBytes(from string, bytes []byte){
	msg := conversationMessage.FromBytes(bytes)


	convMHI.handle(from, msg)
}

func (convMHI *ConversationMessageHandlerImpl)handle(from string, msg *conversationMessage.ConversationMessage){

	decrypted := convMHI.commonKeyProtocol.Decrypt(msg.GetMessageContent())

	switch(msg.GetMessageType()){
	case conversationMessage.DEFAULT:
		convMHI.handleDEFAULT(from, decrypted)
	case conversationMessage.COMMON_KEY_PROTOCOL:
		convMHI.handleCOMMON_KEY_PROTOCOL(decrypted)
	case conversationMessage.INIT_DATA:
		convMHI.handleINIT_DATA(decrypted)
	}
}
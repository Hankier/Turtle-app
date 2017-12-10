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

func (convMHI *ConversationMessageHandlerImpl)HandleBytes(bytes []byte){
	msg := conversationMessage.FromBytes(bytes)

	switch(msg.GetMessageType()){
	case conversationMessage.DEFAULT:
		convMHI.handleDEFAULT(bytes)
	case conversationMessage.COMMON_KEY_PROTOCOL:
		convMHI.handleCOMMON_KEY_PROTOCOL(bytes)
	case conversationMessage.INIT_DATA:
		convMHI.handleINIT_DATA(bytes)
	}
}

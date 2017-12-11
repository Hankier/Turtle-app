package conversationMessageHandler

import (
	"commonKeyProtocol"
	"receiverEncrypter"
	"textReceiver"
	"conversationMessage"
	"log"
)

type ConversationMessageHandlerImpl struct{
	ckp      commonKeyProtocol.CommonKeyProtocol
	rkh      *receiverEncrypter.ReceiverEncrypterImpl
	textrecv textReceiver.TextReceiver
}

func New(ckp commonKeyProtocol.CommonKeyProtocol, rkh *receiverEncrypter.ReceiverEncrypterImpl, textrecv textReceiver.TextReceiver)(*ConversationMessageHandlerImpl){
	convMHI := new(ConversationMessageHandlerImpl)
	convMHI.ckp = ckp
	convMHI.rkh = rkh
	convMHI.textrecv = textrecv
	return convMHI
}

func (convMHI *ConversationMessageHandlerImpl)HandleBytes(from string, bytes []byte){
	msg, err := conversationMessage.FromBytes(bytes)

	if err != nil{
		log.Print(err)
		return
	}
	convMHI.handle(from, msg)
}

func (convMHI *ConversationMessageHandlerImpl)handle(from string, msg *conversationMessage.ConversationMessage){

	decrypted, err := convMHI.ckp.Decrypt(msg.GetEncryptionType(), msg.GetMessageContent())
	if err != nil{
		log.Print(err)
		return
	}

	switch msg.GetMessageType() {
	case conversationMessage.DEFAULT:
		convMHI.handleDEFAULT(from, decrypted)
	case conversationMessage.COMMON_KEY_PROTOCOL:
		convMHI.handleCOMMON_KEY_PROTOCOL(decrypted)
	case conversationMessage.INIT_DATA:
		convMHI.handleINIT_DATA(decrypted)
	}
}
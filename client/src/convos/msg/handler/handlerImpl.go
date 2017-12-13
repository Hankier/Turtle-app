package handler

import (
	"commonKeyProtocol"
	"receiverEncrypter"
	"textReceiver"
	"convos/msg"
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
	message, err := msg.FromBytes(bytes)

	if err != nil{
		log.Print(err)
		return
	}
	convMHI.handle(from, message)
}

func (convMHI *ConversationMessageHandlerImpl)handle(from string, message *msg.ConversationMessage){

	decrypted, err := convMHI.ckp.Decrypt(message.GetEncryptionType(), message.GetMessageContent())
	if err != nil{
		log.Print(err)
		return
	}

	switch message.GetMessageType() {
	case msg.DEFAULT:
		convMHI.handleDEFAULT(from, decrypted)
	case msg.COMMON_KEY_PROTOCOL:
		convMHI.handleCOMMON_KEY_PROTOCOL(decrypted)
	case msg.INIT_DATA:
		convMHI.handleINIT_DATA(decrypted)
	}
}
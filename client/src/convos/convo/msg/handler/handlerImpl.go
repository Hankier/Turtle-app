package handler

import (
	"textReceiver"
	"convos/convo/msg"
	"log"
	"convos/convo/key"
	"convos/convo/encrypter"
)

type HandlerImpl struct{
	commonKey key.CommonKey
	enc       encrypter.Encrypter
	textrecv  textReceiver.TextReceiver
}

func New(commonKey key.CommonKey, enc encrypter.Encrypter, textrecv textReceiver.TextReceiver)(*HandlerImpl){
	convMHI := new(HandlerImpl)
	convMHI.commonKey = commonKey
	convMHI.enc = enc
	convMHI.textrecv = textrecv
	return convMHI
}

func (convMHI *HandlerImpl)HandleBytes(from string, bytes []byte){
	message, err := msg.FromBytes(bytes)

	if err != nil{
		log.Print(err)
		return
	}
	convMHI.handle(from, message)
}

func (convMHI *HandlerImpl)handle(from string, message *msg.ConversationMessage){

	decrypted, err := convMHI.commonKey.Decrypt(message.GetEncryptionType(), message.GetMessageContent())
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
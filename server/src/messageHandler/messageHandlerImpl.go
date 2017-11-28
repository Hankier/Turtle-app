package messageHandler

import (
	"sessionsSender"
	"decrypter"
	"log"
)

type MessageHandlerImpl struct{
	sessSender sessionsSender.SessionsSender
	decrypter decrypter.Decrypter
}

func NewMessageHandlerImpl(sessSender sessionsSender.SessionsSender, decrypter decrypter.Decrypter)(*MessageHandlerImpl){
	mhi := new(MessageHandlerImpl)
	mhi.sessSender = sessSender
	mhi.decrypter = decrypter
	return mhi
}

func (handler *MessageHandlerImpl)HandleBytes(bytes []byte){
	log.Print("Handling bytes " + string(bytes))

	msg := FromBytes(bytes)
	msg.Handle(handler.sessSender, handler.decrypter)
}
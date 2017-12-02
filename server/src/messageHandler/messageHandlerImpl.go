package messageHandler

import (
	"sessionsSender"
	"decrypter"
	_"log"
	"time"
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

func (handler *MessageHandlerImpl)HandleBytes(from string, bytes []byte){
	//log.Print("Handling bytes " + string(bytes))

	msg := FromBytes(from, bytes)

	//TODO remove debug delay
	time.Sleep(time.Second)
	msg.Handle(handler.sessSender, handler.decrypter)
}
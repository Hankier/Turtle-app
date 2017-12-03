package messageHandler

import (
	"sessionSender"
	_"log"
	"time"
	"cryptographer"
)

type MessageHandlerImpl struct{
	sessSender sessionSender.SessionSender
	decrypter cryptographer.Cryptographer
}

func NewMessageHandlerImpl(sessSender sessionSender.SessionSender, decrypter cryptographer.Cryptographer)(*MessageHandlerImpl){
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
	msg.handleMSG(handler.sessSender)
}
package messageHandler

import (
	"sessionsSender"
	"cryptographer"
	_"log"
	"time"
	"message"
)

type MessageHandlerImpl struct{
	sessSender sessionsSender.SessionsSender
	decrypter  cryptographer.Cryptographer
}

func NewMessageHandlerImpl(sessSender sessionsSender.SessionsSender, decrypter cryptographer.Cryptographer)(*MessageHandlerImpl){
	mhi := new(MessageHandlerImpl)
	mhi.sessSender = sessSender
	mhi.decrypter = decrypter
	return mhi
}

func (handler *MessageHandlerImpl)HandleBytes(from string, bytes []byte){
	//log.Print("Handling bytes " + string(bytes))

	msg := message.FromBytes(bytes)

	//TODO remove debug delay
	time.Sleep(time.Second)
	handler.handle(from, msg)
}

func (handler *MessageHandlerImpl)handle(from string, msg *message.Message){
	msg.SetMessageContent(handler.decrypter.Decrypt(msg.GetEncType(), msg.GetMessageContent()))

	switch msg.GetMessageType(){
	case message.MSG:
		handler.handleMSG(from, msg)
		break
	case message.MSG_OK:
		handler.handleMSG_OK(from, msg)
		break
	case message.PING:
		handler.handlePING(from, msg)
		break
	}
}
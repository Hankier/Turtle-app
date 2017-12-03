package messageHandler

import (
	"sessionSender"
	_"log"
	"time"
	"cryptographer"
	"message"
)

type MessageHandlerImpl struct{
	sessSender sessionSender.SessionSender
	convosHandler conversationsHandler
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

	msg := message.FromBytes(bytes)

	//TODO remove debug delay
	time.Sleep(time.Second)
	handler.handle(from, msg)
}

func (handler *MessageHandlerImpl)handle(from string, msg *message.Message){
	msg.SetMessageContent(handler.decrypter.Decrypt(msg.GetEncType(), msg.GetMessageContent()))

	switch msg.GetMessageType(){
	case message.DEFAULT:
		handler.handleDEFAULT(from, msg)
		break
	case message.OK:
		handler.handleOK(from, msg)
		break
	case message.PING:
		handler.handlePING(from, msg)
		break
	}
}
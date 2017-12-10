package messageHandler

import (
	"sessionSender"
	_"log"
	"time"
	"cryptographer"
	"message"
	"conversationsHandler"
	"log"
)

type MessageHandlerImpl struct{
	sessSender    sessionSender.SessionSender
	convosHandler conversationsHandler.ConversationsHandler
	cryptograph   cryptographer.Cryptographer
}

func NewMessageHandlerImpl(sessSender sessionSender.SessionSender, convosHandler conversationsHandler.ConversationsHandler, cryptograph cryptographer.Cryptographer)(*MessageHandlerImpl){
	mhi := new(MessageHandlerImpl)
	mhi.sessSender = sessSender
	mhi.convosHandler = convosHandler
	mhi.cryptograph = cryptograph
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
	decrypted, err := handler.cryptograph.Decrypt(msg.GetEncType(), msg.GetMessageContent())
	if err != nil{
		log.Print(err.Error())
		return
	}
	msg.SetMessageContent(decrypted)

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
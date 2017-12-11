package messageHandler

import (
	_"log"
	"time"
	"cryptographer"
	"message"
	"conversationsHandler"
	"log"
	"session/sender"
)

type MessageHandlerImpl struct{
	ss            sender.Sender
	convoshandler conversationsHandler.ConversationsHandler
	crypt         cryptographer.Cryptographer
}

func New(sessSender sender.Sender, convohandler conversationsHandler.ConversationsHandler, crypt cryptographer.Cryptographer)(*MessageHandlerImpl){
	mhi := new(MessageHandlerImpl)
	mhi.ss = sessSender
	mhi.convoshandler = convohandler
	mhi.crypt = crypt
	return mhi
}

func (handler *MessageHandlerImpl)HandleBytes(from string, bytes []byte){
	//log.Print("Handling bytes " + string(bytes))

	msg, err := message.FromBytes(bytes)

	if err != nil{
		log.Print(err)
		return
	}
	//TODO remove debug delay
	time.Sleep(time.Second)
	handler.handle(from, msg)
}

func (handler *MessageHandlerImpl)handle(from string, msg *message.Message){
	decrypted, err := handler.crypt.Decrypt(msg.GetEncType(), msg.GetMessageContent())
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
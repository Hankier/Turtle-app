package messageHandler

import (
	"sessionsSender"
	"cryptographer"
	_"log"
	"time"
	"message"
	"log"
)

type MessageHandlerImpl struct{
	sessSender  sessionsSender.SessionsSender
	cryptograph cryptographer.Cryptographer
}

func NewMessageHandlerImpl(sessSender sessionsSender.SessionsSender, decrypter cryptographer.Cryptographer)(*MessageHandlerImpl){
	mhi := new(MessageHandlerImpl)
	mhi.sessSender = sessSender
	mhi.cryptograph = decrypter
	return mhi
}

func (handler *MessageHandlerImpl)HandleBytes(from string, bytes []byte){
	msg, err := message.FromBytes(bytes);
	if  err != nil{
		log.Print(err)
		return
	}
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
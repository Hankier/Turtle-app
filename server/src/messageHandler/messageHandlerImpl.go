package messageHandler

import (
	"sessionsSender"
	"decrypter"
)

type MessageHandlerImpl struct{
	sessSender sessionsSender.SessionsSender
	decrypt decrypter.Decrypter
}

func NewMessageHandlerImpl(sessSender sessionsSender.SessionsSender, decrypt decrypter.Decrypter)(*MessageHandlerImpl){
	mhi := new(MessageHandlerImpl)
	mhi.sessSender = sessSender
	mhi.decrypt = decrypt
	return mhi
}

func (*MessageHandlerImpl)HandleBytes(bytes []byte){
	//TODO
}
package parser

import (
	_"log"
	"time"
	"crypt"
	"log"
	"sessions"
	"convos"
	"msgs/msg"
	"msgs/parser/decrypter"
)

type ParserImpl struct{
	sender        sessions.Sender
	receiver      convos.Receiver
	dec 		  crypt.Decrypter
}

func New(sender sessions.Sender, receiver convos.Receiver)(*ParserImpl){
	mhi := new(ParserImpl)
	mhi.sender = sender
	mhi.receiver = receiver
	mhi.dec = decrypter.New()
	return mhi
}

func (pars *ParserImpl)HandleBytes(from string, bytes []byte){
	//log.Print("Handling bytes " + string(bytes))

	message, err := msg.FromBytes(bytes)

	if err != nil{
		log.Print(err)
		return
	}
	//TODO remove debug delay
	time.Sleep(time.Second)
	pars.handle(from, message)
}

func (pars *ParserImpl)handle(from string, message *msg.Message){
	decrypted, err := pars.dec.Decrypt(message.GetEncType(), message.GetMessageContent())
	if err != nil{
		log.Print(err.Error())
		return
	}
	message.SetMessageContent(decrypted)

	switch message.GetMessageType(){
	case msg.DEFAULT:
		pars.handleDEFAULT(from, message)
		break
	case msg.OK:
		pars.handleOK(from, message)
		break
	case msg.PING:
		pars.handlePING(from, message)
		break
	}
}
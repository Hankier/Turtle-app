package parser

import (
	_"log"
	"time"
	"crypt"
	"log"
	"msgs/msg"
	"msgs/parser/decrypter"
	"sessions/sender"
)

type ParserImpl struct{
	sessSender    sender.Sender
	dec 		  crypt.Decrypter
}

func New(sessSender sender.Sender)(Parser){
	mhi := new(ParserImpl)
	mhi.sessSender = sessSender
	mhi.dec = decrypter.New()
	return mhi
}

func (pars *ParserImpl)ParseBytes(from string, bytes []byte){
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
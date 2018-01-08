package parser

import (
	_"log"
	"time"
	"crypt"
	"log"
	"turtleProtocol/msg"
	"server/server/decrypter"
	"server/sessions/sender"
	"math/rand"
)

type ParserImpl struct{
	sessSender    sender.Sender
	dec 		  crypt.Decrypter
}

func New(sessSender sender.Sender)(*ParserImpl){
	mhi := new(ParserImpl)
	mhi.sessSender = sessSender
	mhi.dec = decrypter.New()
	return mhi
}

func (pars *ParserImpl)ParseMessage(from string, message *msg.Message){
	//TODO remove debug delay
	ms := rand.Intn(500) + 250; //random (250,750) ms
	time.Sleep(time.Duration(ms) * time.Millisecond)
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
	case msg.PING:
		pars.handlePING(from, message)
		break
	}
}
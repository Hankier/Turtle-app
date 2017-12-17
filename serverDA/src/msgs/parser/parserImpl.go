package parser

import (
	_"log"
	"time"
	"log"
	"msgs/msg"
	"sessions/sender"
)

type ParserImpl struct{
	sessSender    sender.Sender
}

func New(sessSender sender.Sender)(Parser){
	mhi := new(ParserImpl)
	mhi.sessSender = sessSender
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

	switch message.GetMessageType(){
	case msg.OK:
		pars.handleOK(from, message)
		break
	case msg.ADD:
		pars.handleADD(from, message)
		break
	case msg.DEFAULT:
		pars.handleREMOVE(from, message)
		break
	}
}

package parser

import (
	"log"
	_"serverlist"
	"msgs/msg"
)

func (pars *ParserImpl)handleADD(from string, message *msg.Message){
	log.Print("Received DEFAULT from: " + from)

	pars.sessSender.SendInstant(from, msg.NewMessageOK().ToBytes())

	if len(message.GetMessageContent()) < 8{
		log.Print("Unexpected message end")
		return
	}
	newMsg, err := msg.FromBytes(message.GetMessageContent())
	if err != nil{
		log.Print(err)
		return
	}

	println("newMSG: "+ string(newMsg.GetMessageContent()))
	log.Print("handleMSG msg: ")
}

func (pars *ParserImpl)handleOK(from string, message *msg.Message){
	//TODO
	println("ITS OK")
}

func (pars *ParserImpl)handleREMOVE(from string, message *msg.Message){
	println("//TODO handleREMOVE")
}

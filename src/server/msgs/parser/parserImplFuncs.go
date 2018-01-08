package parser

import (
	"log"
	"turtleProtocol/msg"
)

func (pars *ParserImpl)handleDEFAULT(from string, message *msg.Message){
	log.Print("Received DEFAULT from: " + from)

	if len(message.GetMessageContent()) < 8{
		log.Print("Unexpected message end")
		return
	}
	nextName := string(message.GetMessageContent()[0:8])
	newMsg, err := msg.FromBytes(message.GetMessageContent()[8:])
	if err != nil{
		log.Print(err)
		return
	}

	log.Print("Pushing DEFAULT to: " + nextName)

	pars.sessSender.Send(nextName, newMsg)

	//log.Print("handleMSG, nextName: " + nextName + " msg " + string(bytes))
}

func (pars *ParserImpl)handlePING(from string, message *msg.Message){
	//TODO real PING
	log.Print("RECEIVED PING")
}
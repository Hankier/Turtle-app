package parser

import (
	"log"
	"msgs/msg"
)

func (pars *ParserImpl)handleDEFAULT(from string, message *msg.Message){
	log.Print("Received DEFAULT from: " + from)

	pars.sender.SendInstant(from, msg.NewMessageOK().ToBytes())

	if len(msg.GetMessageContent()) < 8{
		log.Print("Unexpected message end")
		return
	}
	nextName := string(msg.GetMessageContent()[0:8])
	newMsg, err := message.FromBytes(msg.GetMessageContent()[8:])
	if err != nil{
		log.Print(err)
		return
	}

	log.Print("Pushing DEFAULT to: " + nextName)

	handler.sessSender.SendTo(nextName, newMsg)

	//log.Print("handleMSG, nextName: " + nextName + " msg " + string(bytes))
}

func (pars *ParserImpl)handleOK(from string, message *msg.Message){
	pars.sender.UnlockSending(from)
}

func (pars *ParserImpl)handlePING(from string, message *msg.Message){
	pars.sender.SendInstant(from, msg.NewMessageOK().ToBytes())
	//TODO real PING
	log.Print("RECEIVED PING")
}
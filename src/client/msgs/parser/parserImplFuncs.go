package parser

import (
	"log"
	"client/msgs/msg"
)

func (pars *ParserImpl)handleDEFAULT(from string, message *msg.Message){
	//log.Print("DEBUG Received DEFAULT from: " + from)

	pars.sender.SendInstant(from, msg.NewMessageOK().ToBytes())

	if len(message.GetMessageContent()) < 8{
		log.Print("Unexpected message end")
		return
	}
	convoFrom := string(message.GetMessageContent()[0:16])
	content := message.GetMessageContent()[16:]

	//log.Print("DEBUG DEFAULT content: " + (string)(content))

	pars.receiver.OnReceive(convoFrom, content)

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
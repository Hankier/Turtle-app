package parser

import (
	"log"
	"turtleProtocol/msg"
)

func (pars *ParserImpl)handleDEFAULT(from string, message *msg.Message){
	//log.Print("DEBUG Received DEFAULT from: " + from)

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

func (pars *ParserImpl)handlePING(from string, message *msg.Message){
	//TODO real PING
	log.Print("RECEIVED PING")
}
package parser

import (
	"log"
	"message"
)

func (handler *ParserImpl)handleDEFAULT(from string, msg *message.Message){
	//log.Print("DEBUG Received DEFAULT from: " + from)

	handler.ss.SendInstant(message.NewMessageOK())

	if len(msg.GetMessageContent()) < 8{
		log.Print("Unexpected message end")
		return
	}
	receiver := string(msg.GetMessageContent()[0:8])
	receiverServer := string(msg.GetMessageContent()[8:16])
	content := msg.GetMessageContent()[16:]

	//log.Print("DEBUG DEFAULT content: " + (string)(content))

	handler.convoshandler.ReceiveMessage(content, receiver, receiverServer)

	//log.Print("handleMSG, nextName: " + nextName + " msg " + string(bytes))
}

func (handler *ParserImpl)handleOK(from string, msg *message.Message){
	handler.ss.UnlockSending()
}

func (handler *ParserImpl)handlePING(from string, msg *message.Message){
	msgOk := new(message.Message)
	msgOk.SetMessageType(message.OK)
	msgOk.SetMessageContent(make([]byte,0))
	handler.ss.SendInstant(msgOk)

	//TODO real PING
	log.Print("RECEIVED PING")
}
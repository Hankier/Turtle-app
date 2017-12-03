package messageHandler

import (
	"log"
	"message"
)

func (handler *MessageHandlerImpl)handleMSG(from string, msg *message.Message){
	if len(msg.GetMessageContent()) < 8{
		log.Print("Unexpected message end")
		return
	}
	nextName := string(msg.GetMessageContent()[0:8])

	msg.SetMessageContent(append([]byte(nil), msg.GetMessageContent()[8:]...))

	handler.sessSender.SendTo(nextName, msg)

	msgOk := new(message.Message)
	msgOk.SetMessageType(message.MSG_OK)
	msgOk.SetMessageContent(make([]byte,0))

	//log.Print("handleMSG, nextName: " + nextName + " msg " + string(bytes))

	handler.sessSender.SendInstantTo(from, msgOk)
}

func (handler *MessageHandlerImpl)handleMSG_OK(from string, msg *message.Message){
	handler.sessSender.UnlockSending(from)
}

func (handler *MessageHandlerImpl)handlePING(from string, msg *message.Message){
	msgOk := new(message.Message)
	msgOk.SetMessageType(message.MSG_OK)
	msgOk.SetMessageContent(make([]byte,0))
	handler.sessSender.SendInstantTo(from, msgOk)

	//TODO real PING
	log.Print("RECEIVED PING")
}
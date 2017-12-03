package messageHandler

import (
	"log"
	"message"
)

func (handler *MessageHandlerImpl)handleMSG(msg *message.Message){
	if len(msg.GetMessageContent()) < 8{
		log.Print("Unexpected message end")
		return
	}
	convoAndServerName := string(msg.GetMessageContent()[0:16])

	msg.SetMessageContent(append([]byte(nil), msg.GetMessageContent()[16:]...))

	handler.convosHandler.SendTo(convoAndServerName, msg)

	msgOk := new(message.Message)
	msgOk.SetMessageType(message.MSG_OK)
	msgOk.SetMessageContent(make([]byte,0))

	//log.Print("handleMSG, nextName: " + nextName + " msg " + string(bytes))

	handler.sessSender.SendInstantTo(msgOk)
}

func (handler *MessageHandlerImpl)handleMSG_OK(msg *message.Message){
	handler.sessSender.UnlockSending()
}

func (handler *MessageHandlerImpl)handlePING(msg *message.Message){
	msgOk := new(message.Message)
	msgOk.SetMessageType(message.MSG_OK)
	msgOk.SetMessageContent(make([]byte,0))
	handler.sessSender.SendInstantTo(msgOk)

	//TODO real PING
	log.Print("RECEIVED PING")
}
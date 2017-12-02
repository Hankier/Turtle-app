package messageHandler

import (
	"log"
)

func (handler *MessageHandlerImpl)handleMSG(msg *Message){
	if len(msg.messageContent) < 8{
		log.Print("Unexpected message end")
		return
	}
	nextName := string(msg.messageContent[0:8])

	msg.messageContent = append([]byte(nil), msg.messageContent[8:]...)

	bytes := msg.ToBytes()
	handler.sessSender.SendTo(nextName, bytes)

	msgOk := new(Message)
	msgOk.messageType = MSG_OK
	msgOk.messageContent = make([]byte,0)

	//log.Print("handleMSG, nextName: " + nextName + " msg " + string(bytes))

	handler.sessSender.SendInstantTo(msg.previousName, msgOk.ToBytes())
}

func (handler *MessageHandlerImpl)handleMSG_OK(msg *Message){
	handler.sessSender.UnlockSending(msg.previousName)
}

func (handler *MessageHandlerImpl)handlePING(msg *Message){
	msgOk := new(Message)
	msgOk.messageType = MSG_OK
	msgOk.messageContent = make([]byte,0)
	handler.sessSender.SendInstantTo(msg.previousName, msgOk.ToBytes())

	//TODO real PING
	log.Print("RECEIVED PING")
}
package messageHandler

import (
	"log"
	"message"
)

func (handler *MessageHandlerImpl)handleDEFAULT(from string, msg *message.Message){
	log.Print("Received DEFAULT from: " + from)

	//confirm receive
	handler.sessSender.SendInstantTo(from, message.NewMessageOK())

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

func (handler *MessageHandlerImpl)handleOK(from string, msg *message.Message){
	log.Print("Received OK from: " + from)

	handler.sessSender.UnlockSending(from)
}

func (handler *MessageHandlerImpl)handlePING(from string, msg *message.Message){
	log.Print("Received PING from: " + from)

	handler.sessSender.SendInstantTo(from, message.NewMessageOK())

	//TODO real PING
}
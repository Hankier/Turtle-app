package messageHandler

import (
	"log"
	"message"
)

func (handler *MessageHandlerImpl)handleDEFAULT(from string, msg *message.Message){
	log.Print("Received DEFAULT from: " + from)

	if len(msg.GetMessageContent()) < 8{
		log.Print("Unexpected message end")
		return
	}
	nextName := string(msg.GetMessageContent()[0:8])

	msg.SetMessageContent(append([]byte(nil), msg.GetMessageContent()[8:]...))

	log.Print("Pushing DEFAULT to: " + nextName)

	handler.sessSender.SendTo(nextName, msg)

	//log.Print("handleMSG, nextName: " + nextName + " msg " + string(bytes))

	handler.sessSender.SendInstantTo(from, message.NewMessageOK())
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
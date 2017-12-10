package messageHandler

import (
	"log"
	"message"
)

func (handler *MessageHandlerImpl)handleDEFAULT(from string, msg *message.Message){
	if len(msg.GetMessageContent()) < 8{
		log.Print("Unexpected message end")
		return
	}
	receiver := string(msg.GetMessageContent()[0:8])
	receiverServer := string(msg.GetMessageContent()[8:16])
	content := msg.GetMessageContent()[16:]

	handler.convosHandler.ReceiveMessage(content, receiver, receiverServer)

	msgOk := new(message.Message)
	msgOk.SetMessageType(message.OK)
	msgOk.SetMessageContent(make([]byte,0))

	//log.Print("handleMSG, nextName: " + nextName + " msg " + string(bytes))

	handler.sessSender.SendInstant(msgOk)
}

func (handler *MessageHandlerImpl)handleOK(from string, msg *message.Message){
	handler.sessSender.UnlockSending()
}

func (handler *MessageHandlerImpl)handlePING(from string, msg *message.Message){
	msgOk := new(message.Message)
	msgOk.SetMessageType(message.OK)
	msgOk.SetMessageContent(make([]byte,0))
	handler.sessSender.SendInstant(msgOk)

	//TODO real PING
	log.Print("RECEIVED PING")
}
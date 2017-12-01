package messageHandler

import (
	"sessionsSender"
	"log"
)

func (msg *Message)handleMSG(sender sessionsSender.SessionsSender){
	if len(msg.messageContent) < 8{
		log.Print("Unexpected message end")
		return
	}
	nextName := string(msg.messageContent[0:8])

	msg.messageContent = append([]byte(nil), msg.messageContent[8:]...)

	bytes := msg.toBytes()

	sender.SendTo(nextName, bytes)

	msgOk := new(Message)
	msgOk.messageType = MSG_OK
	msgOk.messageContent = make([]byte,0)

	//log.Print("handleMSG, nextName: " + nextName + " msg " + string(bytes))

	sender.SendInstantTo(msg.previousName, msgOk.toBytes())
}

func (msg *Message)handleMSG_OK(sender sessionsSender.SessionsSender){
	sender.UnlockSending(msg.previousName)
}

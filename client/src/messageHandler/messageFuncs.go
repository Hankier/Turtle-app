package messageHandler

import (
	"sessionSender"
	"log"
)

func (msg *Message)handleMSG(sender sessionSender.SessionSender){
	if len(msg.messageContent) < 8{
		log.Print("Unexpected message end")
		return
	}
	msg.messageContent = append([]byte(nil), msg.messageContent[8:]...)

	bytes := msg.ToBytes()

	sender.SendTo(bytes)

	msgOk := new(Message)
	msgOk.messageType = MSG_OK
	msgOk.messageContent = make([]byte,0)

	//log.Print("handleMSG, nextName: " + nextName + " msg " + string(bytes))

	sender.SendInstantTo(msgOk.ToBytes())
}

func (msg *Message)handleMSG_OK(sender sessionSender.SessionSender){
	sender.UnlockSending()
}

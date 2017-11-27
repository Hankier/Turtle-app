package messageHandler

import "sessionsSender"

func (msg *Message)handleMSG(sender *sessionsSender.SessionsSender){
	nextName := string(msg.messageContent[0:8])

	copy(msg.messageContent, msg.messageContent[8:])

	bytes := msg.toBytes()

	(*sender).SendTo(nextName, bytes)

	msgOk := new(Message)
	msgOk.messageType = MSG_OK
	msgOk.messageContent = make([]byte,0)

	(*sender).SendTo(msg.previousName, msgOk.toBytes())
}

func (msg *Message)handleMSG_OK(sender *sessionsSender.SessionsSender){
	(*sender).UnlockSending(msg.previousName)
}

package messageBuilder

import (
	"serverEntry"
	"cryptographer"
	"message"
)

type MessageBuilder struct{
	Path []serverEntry.ServerEntry
	Receiver string
	ReceiverServer string
	Message string
	MsgType message.TYPE
	EncType cryptographer.TYPE
	MyName string
}

func (msgb *MessageBuilder)AddToPath(srve serverEntry.ServerEntry){
	msgb.Path = append(msgb.Path, srve)
}

func (msgb *MessageBuilder) SetReceiver(rcvr string)  {
	msgb.Receiver = rcvr
}

func (msgb *MessageBuilder) SetReceiverServer(rcvrsrv string)  {
	msgb.ReceiverServer = rcvrsrv
}

func(msgb *MessageBuilder) SetMsgType (p message.TYPE){
	msgb.MsgType = p
}

func (msgb *MessageBuilder) SetEncType(p cryptographer.TYPE){
	msgb.EncType = p
}

func (msgb *MessageBuilder) SetMyName(p string)  {
	msgb.MyName = p
}

func (msgb *MessageBuilder)Build()(*message.Message){

	msgContent := ""

	for _, srv := range msgb.Path{
		msgContent += srv.Name
	}

	msgContent += msgb.ReceiverServer
	msgContent += msgb.Receiver
	msgContent += msgb.Message

	msg := message.Message{msgb.MsgType, msgb.MyName, msgb.EncType, []byte(msgContent)}

	return &msg
}
package messageBuilder

import (
	"serverEntry"
	"messageHandler"
	"cryptographer"
)

type MessageBuilder struct{
	Path []serverEntry.ServerEntry
	Receiver string
	ReceiverServer string
	Message string
	MsgType messageHandler.TYPE
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

func(msgb *MessageBuilder) SetMsgType (p messageHandler.TYPE){
	msgb.MsgType = p
}

func (msgb *MessageBuilder) SetEncType(p cryptographer.TYPE){
	msgb.EncType = p
}

func (msgb *MessageBuilder) SetMyName(p string)  {
	msgb.MyName = p
}

func (msgb *MessageBuilder)Build()(*messageHandler.Message){

	msgContent := ""

	for _, srv := range msgb.Path{
		msgContent += srv.Name
	}

	msgContent += msgb.ReceiverServer
	msgContent += msgb.Receiver
	msgContent += msgb.Message

	msg := messageHandler.Message{msgb.MsgType, msgb.MyName, msgb.EncType, []byte(msgContent)}

	return &msg
}
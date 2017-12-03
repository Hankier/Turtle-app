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
	EncTypeServ cryptographer.TYPE
	EncTypeCli	cryptographer.TYPE
	MyName string
	MyServer string
}

func (msgb *MessageBuilder)AddToPath(srve serverEntry.ServerEntry)(*MessageBuilder){
	msgb.Path = append(msgb.Path, srve)
	return msgb
}

func (msgb *MessageBuilder) SetReceiver(rcvr string) (*MessageBuilder) {
	msgb.Receiver = rcvr
	return msgb
}

func (msgb *MessageBuilder) SetReceiverServer(rcvrsrv string) (*MessageBuilder) {
	msgb.ReceiverServer = rcvrsrv
	return msgb
}

func(msgb *MessageBuilder) SetMsgType (p message.TYPE)(*MessageBuilder){
	msgb.MsgType = p
	return msgb
}

func (msgb *MessageBuilder) SetEncTypeServ(p cryptographer.TYPE)(*MessageBuilder){
	msgb.EncTypeServ = p
	return msgb
}

func (msgb *MessageBuilder) SetEncTypeCli(p cryptographer.TYPE)(*MessageBuilder){
	msgb.EncTypeCli = p
	return msgb
}

func (msgb *MessageBuilder) SetMyName(p string)(*MessageBuilder)  {
	msgb.MyName = p
	return msgb
}

func (msgb *MessageBuilder) SetMyServer(p string)(*MessageBuilder){
	msgb.MyServer = p
	return msgb
}

func (msgb *MessageBuilder)Build()(*message.Message){

	msgPieces := make([][]byte, 0,0)
	msgBytes := make([]byte, 0)
	msgContent := ""


	msgContent = msgb.Message + msgContent

	msgContent = msgb.MyServer + msgb.MyName + msgContent

	msgContent = msgb.Receiver


	piece := message.Message{msgb.MsgType, nil, msgb.EncTypeCli, []byte(msgContent)}

	msgPieces[0] = piece.ToBytes()

	msgContent = msgb.Receiver

	piece = message.Message{msgb.MsgType, nil, msgb.EncTypeServ, []byte(msgContent)}

	msgPieces[1] = piece.ToBytes()

	msgContent = msgb.ReceiverServer

	piece = message.Message{msgb.MsgType, nil, msgb.EncTypeCli, []byte(msgContent)}

	msgPieces[2] = piece.ToBytes()

	for i := 3; i < len(msgb.Path); i++{
		msgContent = msgb.Path[i].Name
		piece = message.Message{msgb.MsgType, nil, msgb.EncTypeCli, []byte(msgContent)}

		msgPieces[i] = piece.ToBytes()
	}

	for i := len(msgPieces)-1; i >= 0; i++ {
		msgBytes = append(msgBytes, msgPieces[i]...)
	}

	msg := message.Message{msgb.MsgType, msgb.MyName, msgb.EncTypeServ, msgBytes}

	return &msg
}
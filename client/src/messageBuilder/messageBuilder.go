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
	msgPieces := make([][]byte, len(msgb.Path) + 2)

	msgContent := ([]byte)(msgb.MyServer + msgb.MyName)
	msgContent = append(msgContent, ([]byte)(msgb.Message)...)

	//TODO ENCRYPTION

	piece := message.Message{msgb.MsgType, msgb.EncTypeCli, msgContent}

	msgPieces[0] = ([]byte)(msgb.Receiver)
	msgPieces[0] = append(msgPieces[0], piece.ToBytes()...)

	piece = message.Message{msgb.MsgType, msgb.EncTypeServ, msgPieces[0]}

	msgPieces[1] = ([]byte)(msgb.ReceiverServer)
	msgPieces[1] = append(msgPieces[1], piece.ToBytes()...)

	for i := 0; i < len(msgb.Path); i++{
		piece = message.Message{msgb.MsgType, msgb.EncTypeServ, msgPieces[i+1]}
		msgPieces[i+2] = ([]byte)(msgb.Path[i].Name)
		msgPieces[i+2] = append(msgPieces[i+2], piece.ToBytes()...)
	}

	msg := &message.Message{msgb.MsgType, msgb.EncTypeServ, msgPieces[len(msgb.Path) + 1]};

	return msg
}
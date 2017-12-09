package messageBuilder

import (
	"serverEntry"
	"cryptographer"
	"message"
)

type MessageBuilder struct{
	path []serverEntry.ServerEntry
	receiver string
	receiverServer string
	messageContent []byte
	msgType message.TYPE
	encTypeServ cryptographer.TYPE
	encTypeCli	cryptographer.TYPE
	myName string
	myServer string
}

func NewMessageBuilder(myName string)(*MessageBuilder){
	msgb := new(MessageBuilder)
	msgb.myName = myName
	return msgb
}

func (msgb *MessageBuilder)SetPath(srve []serverEntry.ServerEntry)(*MessageBuilder){
	msgb.path = srve
	return msgb
}

func (msgb *MessageBuilder) SetReceiver(rcvr string) (*MessageBuilder) {
	msgb.receiver = rcvr
	return msgb
}

func (msgb *MessageBuilder) SetReceiverServer(rcvrsrv string) (*MessageBuilder) {
	msgb.receiverServer = rcvrsrv
	return msgb
}

func(msgb *MessageBuilder) SetMsgType (p message.TYPE)(*MessageBuilder){
	msgb.msgType = p
	return msgb
}

func(msgb *MessageBuilder) SetMsgContent (content []byte)(*MessageBuilder){
	msgb.messageContent = content
	return msgb
}

func (msgb *MessageBuilder) SetEncTypeServ(p cryptographer.TYPE)(*MessageBuilder){
	msgb.encTypeServ = p
	return msgb
}

func (msgb *MessageBuilder) SetEncTypeCli(p cryptographer.TYPE)(*MessageBuilder){
	msgb.encTypeCli = p
	return msgb
}

func (msgb *MessageBuilder) SetMyServer(p string)(*MessageBuilder){
	msgb.myServer = p
	return msgb
}

func (msgb *MessageBuilder)Build()(*message.Message){
	msgPieces := make([][]byte, len(msgb.path) + 2)

	msgContent := ([]byte)(msgb.myServer + msgb.myName)
	msgContent = append(msgContent, msgb.messageContent...)

	//TODO ENCRYPTION

	piece := message.Message{msgb.msgType, msgb.encTypeCli, msgContent}

	msgPieces[0] = ([]byte)(msgb.receiver)
	msgPieces[0] = append(msgPieces[0], piece.ToBytes()...)

	piece = message.Message{msgb.msgType, msgb.encTypeServ, msgPieces[0]}

	msgPieces[1] = ([]byte)(msgb.receiverServer)
	msgPieces[1] = append(msgPieces[1], piece.ToBytes()...)

	for i := 0; i < len(msgb.path); i++{
		piece = message.Message{msgb.msgType, msgb.encTypeServ, msgPieces[i+1]}
		msgPieces[i+2] = ([]byte)(msgb.path[i].Name)
		msgPieces[i+2] = append(msgPieces[i+2], piece.ToBytes()...)
	}

	msg := &message.Message{msgb.msgType, msgb.encTypeServ, msgPieces[len(msgb.path) + 1]};

	return msg
}
package messageBuilder

import (
	"serverList"
	"cryptographer"
	"message"
	"errors"
)

type MessageBuilder struct{
	path []string
	receiver string
	receiverServer string
	srvList *serverList.ServerList
	messageContent []byte
	msgType message.TYPE
	encType	cryptographer.TYPE
	myServer string
}

func NewMessageBuilder(sl *serverList.ServerList)(*MessageBuilder){
	msgb := new(MessageBuilder)
	msgb.srvList = sl
	return msgb
}

func (msgb *MessageBuilder)SetMyServer(ms string)(*MessageBuilder){
	msgb.myServer = ms
	return msgb
}

func (msgb *MessageBuilder)SetPath(srve []string)(*MessageBuilder){
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

func (msgb *MessageBuilder) SetEncType(p cryptographer.TYPE)(*MessageBuilder){
	msgb.encType = p
	return msgb
}

func (msgb *MessageBuilder)Build()(*message.Message, error){
	msgPieces := make([][]byte, len(msgb.path) + 2)

	msgContent := ([]byte)(msgb.messageContent)

	var piece message.Message

	piece = message.Message{msgb.msgType, msgb.encType, msgContent}


	msgPieces[0] = ([]byte)(msgb.receiver)
	msgPieces[0] = append(msgPieces[0], piece.ToBytes()...)

	switch(msgb.encType){
	case cryptographer.PLAIN:
		piece = message.Message{msgb.msgType, msgb.encType, cryptographer.EncryptPlain(msgPieces[0])}
	case cryptographer.ELGAMAL:
		piece = message.Message{msgb.msgType, msgb.encType, cryptographer.EncryptElGamal(msgPieces[0])}
	case cryptographer.RSA:
		piece = message.Message{msgb.msgType, msgb.encType, cryptographer.EncryptRSA(msgb.srvList.GetPublicKeyRSA(msgb.receiver), msgPieces[0])}
	default:
		return nil, errors.New("INVALID ENCRYPTION TYPE")
	}


	msgPieces[1] = ([]byte)(msgb.receiverServer)
	msgPieces[1] = append(msgPieces[1], piece.ToBytes()...)

	for i := 0; i < len(msgb.path); i++{
		switch(msgb.encType){
		case cryptographer.PLAIN:
			piece = message.Message{msgb.msgType, msgb.encType, cryptographer.EncryptPlain(msgPieces[i+1])}
		case cryptographer.ELGAMAL:
			piece = message.Message{msgb.msgType, msgb.encType, cryptographer.EncryptElGamal(msgPieces[i+1])}
		case cryptographer.RSA:
			piece = message.Message{msgb.msgType, msgb.encType, cryptographer.EncryptRSA(msgb.srvList.GetPublicKeyRSA(msgb.path[i]), msgPieces[i+1])}
		default:
			return nil, errors.New("INVALID ENCRYPTION TYPE")
		}
		msgPieces[i+2] = ([]byte)(msgb.path[i])
		msgPieces[i+2] = append(msgPieces[i+2], piece.ToBytes()...)
	}

	var msg *message.Message

	switch(msgb.encType){
	case cryptographer.PLAIN:
		msg = &message.Message{msgb.msgType, msgb.encType, cryptographer.EncryptPlain(msgPieces[len(msgb.path) + 1])}
	case cryptographer.ELGAMAL:
		msg = &message.Message{msgb.msgType, msgb.encType, cryptographer.EncryptElGamal(msgPieces[len(msgb.path) + 1])}
	case cryptographer.RSA:
		msg = &message.Message{msgb.msgType, msgb.encType, cryptographer.EncryptRSA(msgb.srvList.GetPublicKeyRSA(msgb.myServer), msgPieces[len(msgb.path) + 1])}
	default:
		return nil, errors.New("INVALID ENCRYPTION TYPE")
	}

	return msg, nil
}
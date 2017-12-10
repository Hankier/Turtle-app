package messageBuilder

import (
	"serverList"
	"cryptographer"
	"message"
	"errors"
)

type MessageBuilder struct{
	path           []string
	receiver       string
	receiverServer string
	srvList        *serverList.ServerList
	convoBuilder   *conversationMessageBuilder.ConversationMessageBuilder
	msgString	   string
	msgType        message.TYPE
	encType        cryptographer.TYPE
	myServer       string
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

func(msgb *MessageBuilder) SetMsgString (content string)(*MessageBuilder){
	msgb.msgString = content
	return msgb
}

func(msgb *MessageBuilder) SetMsgContentBuilder (builder *conversationMessageBuilder.ConversationMessageBuilder)(*MessageBuilder){
	msgb.convoBuilder = builder
	return msgb
}

func (msgb *MessageBuilder) SetEncType(p cryptographer.TYPE)(*MessageBuilder){
	msgb.encType = p
	return msgb
}

func (msgb *MessageBuilder)Build()(*message.Message, error){
	msgPieces := make([][]byte, len(msgb.path) + 2)

	msgb.convoBuilder.ParseString(msgb.msgString)
	msgContent := msgb.convoBuilder.Build()

	var piece *message.Message
	var encElGamal []byte
	var encRSA []byte
	var err error

	piece = message.NewMessage(msgb.msgType, msgb.encType, msgContent)


	msgPieces[0] = ([]byte)(msgb.receiver)
	msgPieces[0] = append(msgPieces[0], piece.ToBytes()...)

	switch(msgb.encType){
	case cryptographer.PLAIN:
		piece = message.NewMessage(msgb.msgType, msgb.encType, cryptographer.EncryptPlain(msgPieces[0]))
	case cryptographer.ELGAMAL:
		encElGamal, err = cryptographer.EncryptElGamal(msgb.srvList.GetPublicKeyElGamal(msgb.receiver), msgPieces[0])
		if err != nil{
			return nil, err
		}
		piece = message.NewMessage(msgb.msgType, msgb.encType, encElGamal)
	case cryptographer.RSA:
		encRSA, err = cryptographer.EncryptRSA(msgb.srvList.GetPublicKeyRSA(msgb.receiver), msgPieces[0])
		if err != nil{
			return nil, err
		}
		piece = message.NewMessage(msgb.msgType, msgb.encType, encRSA)
	default:
		return nil, errors.New("INVALID ENCRYPTION TYPE")
	}


	msgPieces[1] = ([]byte)(msgb.receiverServer)
	msgPieces[1] = append(msgPieces[1], piece.ToBytes()...)

	for i := 0; i < len(msgb.path); i++{
		switch(msgb.encType){
		case cryptographer.PLAIN:
			piece = message.NewMessage(msgb.msgType, msgb.encType, cryptographer.EncryptPlain(msgPieces[i+1]))
		case cryptographer.ELGAMAL:
			encElGamal, _ = cryptographer.EncryptElGamal(msgb.srvList.GetPublicKeyElGamal(msgb.receiver), msgPieces[i+1])
			piece = message.NewMessage(msgb.msgType, msgb.encType, encElGamal)
		case cryptographer.RSA:
			encRSA, err = cryptographer.EncryptRSA(msgb.srvList.GetPublicKeyRSA(msgb.path[i]), msgPieces[i+1])
			if err != nil{
				return nil, err
			}
			piece = message.NewMessage(msgb.msgType, msgb.encType, encRSA)
		default:
			return nil, errors.New("INVALID ENCRYPTION TYPE")
		}
		msgPieces[i+2] = ([]byte)(msgb.path[i])
		msgPieces[i+2] = append(msgPieces[i+2], piece.ToBytes()...)
	}

	var msg *message.Message

	switch(msgb.encType){
	case cryptographer.PLAIN:
		msg = message.NewMessage(msgb.msgType, msgb.encType, cryptographer.EncryptPlain(msgPieces[len(msgb.path) + 1]))
	case cryptographer.ELGAMAL:
		encElGamal, _ = cryptographer.EncryptElGamal(msgb.srvList.GetPublicKeyElGamal(msgb.receiver), msgPieces[len(msgb.path) + 1])
		msg = message.NewMessage(msgb.msgType, msgb.encType, encElGamal)
	case cryptographer.RSA:
		encRSA, err = cryptographer.EncryptRSA(msgb.srvList.GetPublicKeyRSA(msgb.myServer), msgPieces[len(msgb.path) + 1])
		if err != nil{
			return nil, err
		}
		msg = message.NewMessage(msgb.msgType, msgb.encType, encRSA)
	default:
		return nil, errors.New("INVALID ENCRYPTION TYPE")
	}

	return msg, nil
}
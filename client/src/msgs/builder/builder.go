package builder

import (
	"crypt"
	"message"
	"errors"
	"convos/msg/builder"
	"srvlist"
	"msgs/msg"
)

type Builder struct{
	path              []string
	receiver          string
	receiverServer    string
	srvList           *srvlist.ServerList
	convoBuilder      *builder.ConversationMessageBuilder
	receiverEncrypter crypt.Encrypter
	msgString         string
	msgType           msg.TYPE
	encType           crypt.TYPE
	myServer          string
	myName            string
}

func New(myName string, sl *srvlist.ServerList)(*Builder){
	msgb := new(Builder)
	msgb.srvList = sl
	msgb.myName = myName
	return msgb
}

func (msgb *Builder) SetMyCurrentServer(ms string)(*Builder){
	msgb.myServer = ms
	return msgb
}

func (msgb *Builder) SetReceiver(rcvr string) (*Builder) {
	msgb.receiver = rcvr
	return msgb
}

func(msgb *Builder) SetReceiverEncrypter (handler crypt.Encrypter)(*Builder){
	msgb.receiverEncrypter = handler
	return msgb
}

func (msgb *Builder) SetReceiverServer(rcvrsrv string) (*Builder) {
	msgb.receiverServer = rcvrsrv
	return msgb
}

func (msgb *Builder)SetPath(srve []string)(*Builder){
	msgb.path = srve
	return msgb
}

func(msgb *Builder) SetMsgType (p message.TYPE)(*Builder){
	msgb.msgType = p
	return msgb
}

func(msgb *Builder) SetMsgString (content string)(*Builder){
	msgb.msgString = content
	return msgb
}

func(msgb *Builder) SetMsgContentBuilder (builder *builder.ConversationMessageBuilder)(*Builder){
	msgb.convoBuilder = builder
	return msgb
}

func (msgb *Builder) SetEncType(p crypt.TYPE)(*Builder){
	msgb.encType = p
	return msgb
}

func (msgb *Builder)Build()(*message.Message, error){

	if len(msgb.path) > 0{
		if msgb.path[0] == msgb.receiverServer{
			msgb.path = msgb.path[1:]
		}
	}

	if len(msgb.path) > 1{
		if msgb.path[len(msgb.path)-1] == msgb.myServer{
			msgb.path = msgb.path[:len(msgb.path)-1]
		}
	}

	if len(msgb.path) == 1 && (msgb.path[0] == msgb.receiverServer || msgb.path[0] == msgb.myServer){
		msgb.path = make([]string, 0)
	}

	msgPieces := make([][]byte, len(msgb.path) + 2)

	msgb.convoBuilder.ParseCommand(msgb.msgString)
	msgContent := ([]byte)(msgb.myServer)
	msgContent = append(msgContent, ([]byte)(msgb.myName)...)
	msgContent = append(msgContent, ([]byte)(msgb.convoBuilder.Build())...)

	var piece *message.Message

	piece = message.New(msgb.msgType, msgb.encType, msgContent)

	msgPieces[0] = ([]byte)(msgb.receiver)
	msgPieces[0] = append(msgPieces[0], piece.ToBytes()...)


	var encElGamal []byte
	var encRSA []byte
	var err error

	switch(msgb.encType){
	case crypt.PLAIN:
		piece = message.New(msgb.msgType, msgb.encType, msgPieces[0])
	case crypt.ELGAMAL:
		encElGamal, err = msgb.receiverEncrypter.Encrypt(crypt.ELGAMAL, msgPieces[0])
		if err != nil{	return nil, err	}

		piece = message.New(msgb.msgType, msgb.encType, encElGamal)
	case crypt.RSA:
		encRSA, err = msgb.receiverEncrypter.Encrypt(crypt.RSA, msgPieces[0])
		if err != nil{	return nil, err	}

		piece = message.New(msgb.msgType, msgb.encType, encRSA)
	default:
		return nil, errors.New("INVALID ENCRYPTION TYPE")
	}


	msgPieces[1] = ([]byte)(msgb.receiverServer)
	msgPieces[1] = append(msgPieces[1], piece.ToBytes()...)



	var srvEncrypter crypt.Encrypter

	for i := 0; i < len(msgb.path); i++{
		srvEncrypter, err = msgb.srvList.GetEncrypter(msgb.path[i])
		if err != nil{	return nil, err	}

		piece, err = msgb.createPiece(msgPieces[i+1], srvEncrypter)
		if err != nil{	return nil, err	}

		msgPieces[i+2] = ([]byte)(msgb.path[i])
		msgPieces[i+2] = append(msgPieces[i+2], piece.ToBytes()...)
	}

	var msg *message.Message


	srvEncrypter, err = msgb.srvList.GetEncrypter(msgb.myServer)
	if err != nil{	return nil, err	}

	msg, err = msgb.createPiece(msgPieces[len(msgb.path) + 1], srvEncrypter)
	if err != nil{	return nil, err	}

	return msg, nil
}

func (msgb *Builder)createPiece(pieceContent []byte, enc crypt.Encrypter)(*message.Message, error){
	var piece *message.Message
	var encElGamal []byte
	var encRSA []byte
	var err error

	switch msgb.encType {
	case crypt.PLAIN:
		piece = message.New(msgb.msgType, msgb.encType, pieceContent)
	case crypt.ELGAMAL:
		encElGamal, err = enc.Encrypt(crypt.ELGAMAL, pieceContent)
		if err != nil{	return nil, err	}

		piece = message.New(msgb.msgType, msgb.encType, encElGamal)
	case crypt.RSA:
		encRSA, err = enc.Encrypt(crypt.RSA, pieceContent)
		if err != nil{	return nil, err	}

		piece = message.New(msgb.msgType, msgb.encType, encRSA)
	default:
		return nil, errors.New("INVALID ENCRYPTION TYPE")
	}

	return piece, nil
}
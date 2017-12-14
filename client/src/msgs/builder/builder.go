package builder

import (
	"crypt"
	"errors"
	"srvlist"
	"msgs/msg"
	"client"
)

type Builder struct{
	credHolder client.CredentialsHolder
	srvList           *srvlist.ServerList

	path              []string
	receiver          string
	receiverServer    string
	receiverEncrypter crypt.Encrypter
	msgString         string
	msgType           msg.TYPE
	encType           crypt.TYPE
	content			  []byte
}

func New(credHolder client.CredentialsHolder, sl *srvlist.ServerList)(*Builder){
	msgb := new(Builder)
	msgb.srvList = sl
	msgb.credHolder = credHolder
	return msgb
}

func (msgb *Builder) SetReceiver(rcvr string) (*Builder) {
	msgb.receiver = rcvr
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

func(msgb *Builder) SetMsgType (p msg.TYPE)(*Builder){
	msgb.msgType = p
	return msgb
}

func(msgb *Builder) SetMsgString (content string)(*Builder){
	msgb.msgString = content
	return msgb
}

func(msgb *Builder) SetMsgContent (content []byte)(*Builder){
	msgb.content = content
	return msgb
}

func (msgb *Builder) SetEncType(p crypt.TYPE)(*Builder){
	msgb.encType = p
	return msgb
}

func (msgb *Builder)Build()(*msg.Message, error){

	myServer := msgb.credHolder.GetCurrentServer()
	myName := msgb.credHolder.GetName()

	if len(msgb.path) > 0{
		if msgb.path[0] == msgb.receiverServer{
			msgb.path = msgb.path[1:]
		}
	}

	if len(msgb.path) > 1{
		if msgb.path[len(msgb.path)-1] == myServer{
			msgb.path = msgb.path[:len(msgb.path)-1]
		}
	}

	if len(msgb.path) == 1 && (msgb.path[0] == msgb.receiverServer || msgb.path[0] == myServer){
		msgb.path = make([]string, 0)
	}

	msgPieces := make([][]byte, len(msgb.path) + 2)

	msgContent := ([]byte)(myServer)
	msgContent = append(msgContent, ([]byte)(myName)...)
	msgContent = append(msgContent, ([]byte)(msgb.content)...)

	var piece *msg.Message

	piece = msg.New(msgb.msgType, msgb.encType, msgContent)

	msgPieces[0] = ([]byte)(msgb.receiver)
	msgPieces[0] = append(msgPieces[0], piece.ToBytes()...)


	var encElGamal []byte
	var encRSA []byte
	var err error

	switch(msgb.encType){
	case crypt.PLAIN:
		piece = msg.New(msgb.msgType, msgb.encType, msgPieces[0])
	case crypt.ELGAMAL:
		encElGamal, err = msgb.receiverEncrypter.Encrypt(crypt.ELGAMAL, msgPieces[0])
		if err != nil{	return nil, err	}

		piece = msg.New(msgb.msgType, msgb.encType, encElGamal)
	case crypt.RSA:
		encRSA, err = msgb.receiverEncrypter.Encrypt(crypt.RSA, msgPieces[0])
		if err != nil{	return nil, err	}

		piece = msg.New(msgb.msgType, msgb.encType, encRSA)
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

	var message *msg.Message


	srvEncrypter, err = msgb.srvList.GetEncrypter(myServer)
	if err != nil{	return nil, err	}

	message, err = msgb.createPiece(msgPieces[len(msgb.path) + 1], srvEncrypter)
	if err != nil{	return nil, err	}

	return message, nil
}

func (msgb *Builder)createPiece(pieceContent []byte, enc crypt.Encrypter)(*msg.Message, error){
	var piece *msg.Message
	var encElGamal []byte
	var encRSA []byte
	var err error

	switch msgb.encType {
	case crypt.PLAIN:
		piece = msg.New(msgb.msgType, msgb.encType, pieceContent)
	case crypt.ELGAMAL:
		encElGamal, err = enc.Encrypt(crypt.ELGAMAL, pieceContent)
		if err != nil{	return nil, err	}

		piece = msg.New(msgb.msgType, msgb.encType, encElGamal)
	case crypt.RSA:
		encRSA, err = enc.Encrypt(crypt.RSA, pieceContent)
		if err != nil{	return nil, err	}

		piece = msg.New(msgb.msgType, msgb.encType, encRSA)
	default:
		return nil, errors.New("INVALID ENCRYPTION TYPE")
	}

	return piece, nil
}
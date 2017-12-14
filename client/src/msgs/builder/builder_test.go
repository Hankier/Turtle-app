package builder

import (
	"testing"
	"crypt"
	"fmt"
	"bytes"
	"msgs/msg"
	"srvlist"
	"client/credentials"
	"convos/msgsBuilder"
)

type MockCredsHandler struct{
	myname string
	myserver string
}

func (mch *MockCredsHandler)GetName()string{
	return mch.myname
}
func (mch *MockCredsHandler)GetCurrentServer()(string, error){
	return mch.myserver, nil
}

func NewMCH(name, serv string)(*MockCredsHandler){
	mch := new(MockCredsHandler)
	mch.myname = name
	mch.myserver = serv
	return mch
}


type MockConvMsgBuilder struct{
}

func NewMCMB()(*MockConvMsgBuilder){
	mcmb := new(MockConvMsgBuilder)
	return mcmb
}

func (mcmb *MockConvMsgBuilder)BuildMessageContent(server string, name string, command string, encType crypt.TYPE)([]byte, error){
	ret := []byte(server + name)
	ret = append(ret, 0)
	ret = append(ret, 0)
	ret = append(ret, []byte(command)...)
	return ret, nil
}

func TestMessageBuilder_Build(t *testing.T) {
	var mch credentials.CredentialsHolder
	mch = NewMCH("10000000", "00000000")

	var mcmb msgsBuilder.MessageBuilder
	mcmb = NewMCMB()

	msgb := New(srvlist.New(), mcmb, mch)

	expected := ([]byte)("  00000002  00000001  recvserv  recvrecv  0000000010000000  abcd")
	expected[0] = 0
	expected[1] = 0
	expected[10] = 0
	expected[11] = 0
	expected[20] = 0
	expected[21] = 0
	expected[30] = 0
	expected[31] = 0
	expected[40] = 0
	expected[41] = 0
	expected[58] = 0
	expected[59] = 0

	msgString := "abcd"

	path := []string{"recvserv", "00000001", "00000002", "00000000"}

	//convoBuilder := builder.New(&commonKeyProtocol.CommonKeyProtocolImpl{})

	msg,_ :=
		msgb.SetMsgContent([]byte(msgString)).
		SetReceiver("recvrecv").
		SetReceiverServer("recvserv").
		SetEncType(crypt.PLAIN).
		SetMsgType(msg.DEFAULT).
		SetPath(path).
		SetCommand(msgString).
		Build()
	fmt.Println(string(msg.ToBytes()))
	fmt.Println(string(expected))

	if !bytes.Equal(msg.ToBytes(), ([]byte)(expected)){
		t.Error("Unexpected message")
	}
}

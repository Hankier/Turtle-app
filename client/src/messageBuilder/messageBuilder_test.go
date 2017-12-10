package messageBuilder

import (
	"testing"
	"serverList"
	"cryptographer"
	"message"
	"fmt"
	"bytes"
	"conversationMessageBuilder"
	"commonKeyProtocol"
	"receiverKeyHandler"
)

func TestMessageBuilder_Build(t *testing.T) {
	msgb := NewMessageBuilder(serverList.NewServerList())
	msgb.SetMyName("10000000").SetMyServer("00000000")

	expected := ([]byte)("H   00000001<   000000020   00000000$   50000000   0000000010000000  abcd")
	expected[1] = 0
	expected[2] = 0
	expected[3] = 0
	expected[3] = 0
	expected[13] = 0
	expected[14] = 0
	expected[15] = 0
	expected[25] = 0
	expected[26] = 0
	expected[27] = 0
	expected[37] = 0
	expected[38] = 0
	expected[39] = 0
	expected[49] = 0
	expected[50] = 0
	expected[51] = 0
	expected[68] = 0
	expected[69] = 0

	msgString := "abcd"

	path := []string{"00000002", "00000001"}

	convoBuilder := conversationMessageBuilder.NewConversationMessageBuilder(&commonKeyProtocol.CommonKeyProtocolImpl{})

	msg,_ :=
		msgb.SetMsgString(msgString).
		SetMsgContentBuilder(convoBuilder).
		SetReceiverKeyHandler(&receiverKeyHandler.ReceiverKeyHandlerImpl{}).
		SetReceiver("50000000").
		SetReceiverServer("00000000").
		SetEncType(cryptographer.PLAIN).
		SetMsgType(message.DEFAULT).
		SetPath(path).
		Build()
	fmt.Println(string(msg.ToBytes()))
	fmt.Println(string(expected))

	if !bytes.Equal(msg.ToBytes(), ([]byte)(expected)){
		t.Error("Unexpected message")
	}
}

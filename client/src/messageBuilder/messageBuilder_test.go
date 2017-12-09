package messageBuilder

import (
	"testing"
	"serverList"
	"cryptographer"
	"message"
	"fmt"
	"bytes"
)

func TestMessageBuilder_Build(t *testing.T) {
	msgb := NewMessageBuilder(serverList.NewServerList())

	comparer := message.NewMessage(message.DEFAULT, cryptographer.PLAIN, []byte("00000000")).ToBytes()
	comparer = append(comparer, message.NewMessage(message.DEFAULT, cryptographer.PLAIN, []byte("00000001")).ToBytes()...)
	comparer = append(comparer, message.NewMessage(message.DEFAULT, cryptographer.PLAIN, []byte("00000002")).ToBytes()...)
	comparer = append(comparer, message.NewMessage(message.DEFAULT, cryptographer.PLAIN, []byte("00000000")).ToBytes()...)
	comparer = append(comparer, message.NewMessage(message.DEFAULT, cryptographer.PLAIN, []byte("50000000")).ToBytes()...)

	comparer = append(comparer, message.NewMessage(message.DEFAULT, cryptographer.PLAIN, []byte("abcd")).ToBytes()...)
	msgContent := "abcd"

	path := []string{"00000002", "00000001", "00000000"}

	msg,_ := msgb.SetEncType(cryptographer.PLAIN).
		SetMsgContent([]byte(msgContent)).
			SetMyServer("00000000").
				SetMsgType(message.DEFAULT).
					SetPath(path).
						SetReceiver("50000000").
							SetReceiverServer("00000000").
								Build()
	fmt.Println(string(msg.ToBytes()))
	fmt.Println(string(comparer))

	if !bytes.Equal(msg.ToBytes(), comparer){
		t.Error("Unexpected message")
	}
}

package parser

import (
	"testing"
	"msgs/msg"
	"crypt"
	"fmt"
	"client/decrypter"
	"bytes"
)

type SenderMock struct{
}

func (*SenderMock)Send(name string, content []byte)error{
	return nil
}
func (*SenderMock)SendInstant(name string, content []byte)error{
	return nil
}
func (*SenderMock)UnlockSending(name string)error{
	return nil
}

type ReceiverMock struct{
	from	string
	handled []byte
}

func (rm *ReceiverMock)	OnReceive(from string, content []byte){
	rm.from = from
	rm.handled = content
}




func TestParserImpl_ParseBytes(t *testing.T) {
	sendermock := new(SenderMock)
	receivermock := new(ReceiverMock)
	parser := New(sendermock, receivermock)

	content := []byte("0000000000000000thisissomemessagemate")

	dec := decrypter.New()

	enccontent, err := crypt.EncryptRSA(dec.GetPublicKey(), content)

	fmt.Println(string(enccontent))

	if err != nil{
		t.Fail()
	}

	message := msg.New(msg.DEFAULT, crypt.RSA, enccontent)

	parser.ParseBytes("0000000000000000", message.ToBytes())

	fmt.Println(receivermock.from + string(receivermock.handled))

	actual := append([]byte(receivermock.from), receivermock.handled...)

	if !bytes.Equal(content, actual) && len(content) == len(actual){
		t.Fail()
	}
}
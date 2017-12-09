package messageBuilder

import (
	"testing"
	"serverList"
	"cryptographer"
	"message"
)

func TestMessageBuilder_Build(t *testing.T) {
	msgb := NewMessageBuilder(serverList.NewServerList())

	comparer := message.Message{message.DEFAULT, cryptographer.PLAIN, []byte("00000000")}


	msgContent := "abcd"

	msgb.SetEncType(cryptographer.PLAIN).SetMsgContent()
}

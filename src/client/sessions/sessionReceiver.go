package sessions

import (
	"client/msgs/parser"
	"client/sessions/sender"
	"client/convos/receiver"
)

type SessionReceiver struct{
	name string
	msgsParser *parser.ParserImpl

}

func NewSessionReceiver(name string, sessionsSender sender.Sender, convosRecv receiver.Receiver) *SessionReceiver{
	r := new(SessionReceiver)
	r.msgsParser = parser.New(sessionsSender, convosRecv)

	return r
}

func (r* SessionReceiver)OnReceive(content []byte){
	r.msgsParser.ParseBytes(r.name, content)
}

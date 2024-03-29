package sessions

import (
	"client/msgs/parser"
	"client/sessions/sender"
	"client/convos/receiver"
	"turtleProtocol/msg"
)

type SessionReceiver struct{
	name string
	msgsParser *parser.ParserImpl

}

func NewSessionReceiver(name string, sessionsSender sender.Sender, convosRecv receiver.Receiver) *SessionReceiver{
	r := new(SessionReceiver)
	r.msgsParser = parser.New(sessionsSender, convosRecv)
	r.name = name

	return r
}

func (r* SessionReceiver)OnReceive(message *msg.Message){
	r.msgsParser.ParseMessage(r.name, message)
}

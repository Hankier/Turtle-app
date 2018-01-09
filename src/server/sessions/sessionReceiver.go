package sessions

import (
	"server/msgs/parser"
	"server/sessions/sender"
	"turtleProtocol/msg"
)

type SessionReceiver struct{
	name string
	msgsParser *parser.ParserImpl

}

func NewSessionReceiver(name string, sessionsSender sender.Sender) *SessionReceiver{
	r := new(SessionReceiver)
	r.msgsParser = parser.New(sessionsSender)
	r.name = name

	return r
}

func (r* SessionReceiver)OnReceive(message *msg.Message){
	r.msgsParser.ParseMessage(r.name, message)
}

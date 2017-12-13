package sessions

import (
	"sessions/session"
	"msgs/parser"
	"net"
	"convos"
)

type Controller struct{
	sessions   map[string]*session.Session
	convosRecver convos.Receiver
	msgsParser parser.Parser
}

func New(msgsParser parser.Parser)(*Controller){
	c := new(Controller)
	c.sessions = make(map[string]*session.Session)

	c.msgsParser = msgsParser
	return c
}

func (c *Controller)CreateSession(name string, socket net.Conn){
	sess := session.New(socket, name, c, c)

	go sess.Start()

	c.sessions[name] = sess
	//TODO thread safe
}

func (c *Controller)RemoveSession(name string){
	c.sessions[name].DeleteSession()
	delete(c.sessions, name)
}

func (c *Controller)GetActiveSessions()[]string{
	//TODO

	return nil
}

func (c *Controller)OnReceive(name string, content []byte){
	c.msgsParser.ParseBytes(name, content)
}


package sessions

import (
	"sessions/session"
	"msgs/parser"
	"net"
	"convos"
	"errors"
)

type Controller struct{
	sessions   map[string]*session.Session
	msgsParser parser.Parser
}

func New(convosRecver convos.Receiver)(*Controller){
	c := new(Controller)
	c.sessions = make(map[string]*session.Session)
	c.msgsParser = parser.New(c, convosRecver)
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

func (c *Controller)SendTo(name string, content []byte)error{
	if sess, ok := c.sessions[name]; ok {
		sess.Send(content)
	}else{
		return errors.New("wrong session name")
	}
	return nil
}

func (c *Controller)SendInstantTo(name string, content []byte)error{
	if sess, ok := c.sessions[name]; ok {
		sess.SendInstant(content)
	}else{
		return errors.New("wrong session name")
	}
	return nil
}

func (c *Controller)UnlockSending(name string)error{
	if sess, ok := c.sessions[name]; ok {
		sess.UnlockSending()
	}else{
		return errors.New("wrong session name")
	}
	return nil
}


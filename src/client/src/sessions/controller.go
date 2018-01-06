package sessions

import (
	"sessions/session"
	"msgs/parser"
	"net"
	"errors"
	"convos/receiver"
	"sync"
	"log"
	"reflect"
)

type Controller struct{
	sessions struct{
		sync.Mutex
		data map[string]*session.Session
	}
	msgsParser parser.Parser
}

func New(convosRecver receiver.Receiver)(*Controller){
	c := new(Controller)
	c.sessions.data = make(map[string]*session.Session)
	c.msgsParser = parser.New(c, convosRecver)
	return c
}

func (c *Controller)CreateSession(name string, socket net.Conn){
	sess := session.New(socket, name, c, c)

	go sess.Start()

	c.sessions.Lock()
	c.sessions.data[name] = sess
	c.sessions.Unlock()
	log.Print("Creating session with: " + name)
}

func (c *Controller)RemoveSession(name string){
	c.sessions.Lock()
	c.sessions.data[name].DeleteSession()
	delete(c.sessions.data, name)
	c.sessions.Unlock()
	log.Print("Removing session with: " + name)
}

func (c *Controller)GetActiveSessions()[]string{
	c.sessions.Lock()

	activeSessions := make([]string, len(c.sessions.data))
	i := 0
	for name, _ := range c.sessions.data{
		activeSessions[i] = name
		i++
	}

	c.sessions.Unlock()
	return activeSessions
}

func (c *Controller)OnReceive(from string, content []byte){
	c.msgsParser.ParseBytes(from, content)
}

func (c *Controller)Send(name string, content []byte)error{
	c.sessions.Lock()
	sess, ok := c.sessions.data[name]
	c.sessions.Unlock()

	if ok {
		sess.Send(content)
	}else{
		return errors.New(reflect.TypeOf(c).String() + " wrong session name")
	}


	return nil
}

func (c *Controller)SendInstant(name string, content []byte)error{
	c.sessions.Lock()
	sess, ok := c.sessions.data[name]
	c.sessions.Unlock()

	if  ok {
		sess.SendInstant(content)
	}else{
		return errors.New(reflect.TypeOf(c).String() + " wrong session name")
	}

	return nil
}

func (c *Controller)UnlockSending(name string)error{
	c.sessions.Lock()
	sess, ok := c.sessions.data[name];
	c.sessions.Unlock()

	if  ok {
		sess.UnlockSending()
	}else{
		return errors.New(reflect.TypeOf(c).String() +" wrong session name")
	}

	return nil
}


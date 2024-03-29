package sessions

import (
	"net"
	"errors"
	convosReceiver "client/convos/receiver"
	"sync"
	"log"
	"reflect"
	"turtleProtocol"
	"turtleProtocol/msg"
)

type Controller struct{
	sessions struct{
		sync.Mutex
		sess map[string]*turtleProtocol.Session
		recv map[string]*SessionReceiver
	}
	convosRecv convosReceiver.Receiver
}

func New(convosRecv convosReceiver.Receiver)(*Controller){
	c := new(Controller)
	c.sessions.sess = make(map[string]*turtleProtocol.Session)
	c.sessions.recv = make(map[string]*SessionReceiver)
	c.convosRecv = convosRecv
	return c
}

func (c *Controller)CreateSession(name string, socket net.Conn){
	recv := NewSessionReceiver(name, c, c.convosRecv)
	sess := turtleProtocol.NewSession(socket, name, recv)

	c.sessions.Lock()
	c.sessions.sess[name] = sess
	c.sessions.recv[name] = recv
	c.sessions.Unlock()

	go c.startSession(sess)

	log.Print("Creating session with: " + name)
}

func (c* Controller)startSession(session* turtleProtocol.Session){
	session.Start()
	c.RemoveSession(session.GetName())
}

func (c *Controller)RemoveSession(name string){
	c.sessions.Lock()
	sess, ok := c.sessions.sess[name]
	//session not already deleted
	if ok {
		sess.DeleteSession()
		delete(c.sessions.sess, name)
		delete(c.sessions.recv, name)
	}
	c.sessions.Unlock()
	log.Print("Removing session with: " + name)
}

func (c *Controller)GetActiveSessions()[]string{
	c.sessions.Lock()

	activeSessions := make([]string, len(c.sessions.sess))
	i := 0
	for name, _ := range c.sessions.sess{
		activeSessions[i] = name
		i++
	}

	c.sessions.Unlock()
	return activeSessions
}

func (c *Controller)Send(name string, message *msg.Message)error{
	c.sessions.Lock()
	sess, ok := c.sessions.sess[name]
	c.sessions.Unlock()

	if ok {
		sess.Send(message)
	} else {
		return errors.New(reflect.TypeOf(c).String() + " wrong session name")
	}

	return nil
}


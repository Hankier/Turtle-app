package sessions

import (
	"server/msgs/parser"
	"net"
	"errors"
	"sync"
	"log"
	"server/server/dialer"
	"server/server/credentials"
	"turtleProtocol"
	"turtleProtocol/msg"
)

type Controller struct{
	sessions struct{
		sync.Mutex
		sess map[string]*turtleProtocol.Session
		recv map[string]*SessionReceiver
	}
	msgsParser parser.Parser
	serverDialer dialer.Dialer
	credHolder credentials.CredentialsHolder
}

func New(serverDialer dialer.Dialer, credHolder credentials.CredentialsHolder)(*Controller){
	c := new(Controller)
	c.sessions.sess = make(map[string]*turtleProtocol.Session)
	c.sessions.recv = make(map[string]*SessionReceiver)

	c.msgsParser = parser.New(c)
	c.serverDialer = serverDialer
	c.credHolder = credHolder

	return c
}

func (c *Controller)CreateSession(name string, socket net.Conn){
	recv := NewSessionReceiver(name, c)
	sess := turtleProtocol.NewSession(socket, name, recv)

	go c.startSession(sess)

	c.sessions.Lock()
	c.sessions.sess[name] = sess
	c.sessions.recv[name] = recv
	c.sessions.Unlock()
	log.Print("Creating session with: " + name)
}


func (c* Controller)startSession(session* turtleProtocol.Session){
	session.Start()
	c.RemoveSession(session.GetName())
}

func (c *Controller)RemoveSession(name string){
	c.sessions.Lock()
	sess, ok := c.sessions.sess[name]
	//not already removed
	if ok{
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

func (c *Controller)OnReceive(from string, message *msg.Message){
	c.msgsParser.ParseMessage(from, message)
}

func (c *Controller)Send(name string, message *msg.Message)error{
	//sending to myself
	if name == c.credHolder.GetName(){
		c.OnReceive(name, message)
		return nil
	}

	c.sessions.Lock()
	sess, ok := c.sessions.sess[name]
	c.sessions.Unlock()

	if ok {
		sess.Send(message)
	}else{
		err := c.serverDialer.ConnectToServer(name)
		if err != nil {
			return errors.New("wrong session name")
		} else {
			c.sessions.Lock()
			sess, ok = c.sessions.sess[name]
			c.sessions.Unlock()
			if ok {
				sess.Send(message)
			} else {
				return errors.New("wrong session name")
			}
		}
	}

	return nil
}


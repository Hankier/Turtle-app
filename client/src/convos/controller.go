package convos

import (
	"convos/convo"
	"sync"
	"errors"
	"textReceiver"
)

type Controller struct{
	conversations struct{
		sync.Mutex
		data map[string]*convo.Conversation
	}

	textRecver textReceiver.TextReceiver
}

func New(textRecver textReceiver.TextReceiver)(*Controller){
	c := new(Controller)
	c.conversations.data = make(map[string]*convo.Conversation)
	c.textRecver = textRecver

	return c
}

func (c *Controller)CreateConversation(server string, name string)(err error){
	convoname := server + name

	c.conversations.Lock()
	conv, ok := c.conversations.data[convoname]
	if !ok{
		conv = convo.New(c.textRecver, server, name)
		c.conversations.data[convoname] = conv
	} else {
		err = errors.New("conversation already exists")
	}
	c.conversations.Unlock()
	return err
}

func (c *Controller)RemoveConversation(server string, name string){
	//TODO REMOVING CONVO WITH ERROR CHECKING
}

func (c *Controller)GetActiveConversations()[]*struct{
	server string
	name string
}{
	//TODO
	return nil
}

func (c *Controller)OnReceive(server string, name string, content []byte){
	convoname := server + name

	//TODO CHECK IF EXISTS
	c.conversations.data[convoname].Receive(content)
}
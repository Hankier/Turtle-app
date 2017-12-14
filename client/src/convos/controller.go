package convos

import (
	"convos/convo"
	"sync"
	"errors"
	"textReceiver"
	"crypt"
	"client/credentials"
)

type Controller struct{
	conversations struct{
		sync.Mutex
		data map[string]*convo.Conversation
	}

	textRecver textReceiver.TextReceiver
	credHolder credentials.CredentialsHolder
}

func New(textRecver textReceiver.TextReceiver, credHolder credentials.CredentialsHolder)(*Controller){
	c := new(Controller)
	c.conversations.data = make(map[string]*convo.Conversation)
	c.textRecver = textRecver
	c.credHolder = credHolder
	return c
}

func (c *Controller)CreateConversation(server string, name string)(err error){
	convoname := server + name

	c.conversations.Lock()
	conv, ok := c.conversations.data[convoname]
	if !ok{
		conv = convo.New(server, name, c.textRecver, c.credHolder)
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

func (c *Controller)OnReceive(from string, content []byte){
	server := from[0:8]
	name := from[8:16]
	conv, ok := c.conversations.data[from]
	if !ok{
		c.CreateConversation(server, name)
		conv, _ = c.conversations.data[from]
	}
	conv.Receive(content)
}

func (c *Controller)BuildMessageContent(server string, name string, command string, encType crypt.TYPE)([]byte, error){
	convoname := server + name

	var content []byte
	var err error

	conv, ok := c.conversations.data[convoname]
	if !ok{
		c.CreateConversation(server, name)
		conv, _ = c.conversations.data[convoname]
	}
	content = conv.BuildMessageContent(command, encType)
	return content, err
}
package client

import (
	"session"
	"srvlist"
	"log"
	"net"
	"messageHandler"
	"crypt"
	"errors"
	"message"
	"textReceiver"
	"conversation"
	"messageBuilder"
	"sync"
	"commandsListener"
)

type Client struct{
	srvList        *srvlist.ServerList
	sess           *session.Session
	cmdListener    *commandsListener.CommandsListener
	nodeCrypto     crypt.Cryptographer
	currentPath    []string
	convosMutex    sync.Mutex
	conversations  map[string]*conversation.Conversation
	messageBuilder *messageBuilder.MessageBuilder
	textReceiver   textReceiver.TextReceiver
	myName         string
}

func NewClient(name string)(*Client){
	cli := new(Client)

	cli.myName = name
	cli.srvList = srvlist.New()
	cli.nodeCrypto = crypt.New()
	cli.textReceiver = &textReceiver.TextReceiverImpl{}
	cli.messageBuilder = messageBuilder.NewMessageBuilder(cli.srvList)
	cli.messageBuilder.SetMyName(cli.myName)
	cli.cmdListener = commandsListener.New(cli, cli.textReceiver)
	cli.conversations = make(map[string]*conversation.Conversation)

	return cli
}

func (cli *Client)Start(){
	cli.cmdListener.Listen()
}

func (cli *Client)CreateSession(name string, socket net.Conn){
	if cli.sess != nil{
		cli.RemoveSession()
	}
	msgHandler := messageHandler.New(cli, cli, cli.nodeCrypto)

	sess := session.New(socket, name, msgHandler, cli)

	go sess.Start()
	cli.sess = sess
	//TODO thread safe
}

func (cli *Client)RemoveSession(){
	cli.sess.DeleteSession()
	cli.sess = nil
}

func (cli *Client)Send(msg *message.Message)error{
	if cli.sess != nil{
		cli.sess.Send(msg)
		return nil
	}else{
		log.Println("Not connected to any server\n");
		return errors.New("NOT CONNECTED")
	}
}

func (cli *Client)SendInstant(msg *message.Message)error{
	if cli.sess != nil{
		cli.sess.SendInstant(msg)
		return nil
	}else{
		log.Println("Not connected to any server\n");
		return errors.New("NOT CONNECTED")
	}
}

func (cli *Client)UnlockSending(){
	cli.sess.UnlockSending()
}

func (cli *Client)GetCurrentPath() []string{
	return cli.currentPath
}

func (cli *Client)ChooseNewPath(length int)([]string, error){
	var err error
	cli.currentPath, err = cli.srvList.GetRandomPath(length)

	if err != nil{
		return nil, err
	}

	return cli.currentPath, nil
}

func (cli *Client)ConnectToServer(name string)error{
	socket, err := net.Dial("tcp", cli.srvList.GetServerIpPort(name))
	if err != nil {
		return err
	}
	socket.Write([]byte(cli.myName))
	cli.messageBuilder.SetMyServer(name)
	cli.CreateSession(name, socket)
	log.Print("Succesfully connected to " + name)
	return nil
}

func (cli *Client)GetServerList()[]string{
	return cli.srvList.GetServerList()
}

func (cli *Client)CreateConversation(receiver string, receiverServer string) (convo *conversation.Conversation, err error){
	name := receiverServer + receiver

	cli.convosMutex.Lock()
	convo, ok := cli.conversations[name]
	if !ok{
		convo = conversation.New(cli.textReceiver, receiver, receiverServer)
		cli.conversations[name] = convo
	} else {
		err = errors.New("conversation already exists")
	}
	cli.convosMutex.Unlock()
	return convo, err
}

func (cli *Client)SendTo(message string, receiver string, receiverServer string)error{
	name := receiverServer + receiver

	cli.convosMutex.Lock()
	convo, ok := cli.conversations[name]
	cli.convosMutex.Unlock()
	if !ok{
		newConvo, err := cli.CreateConversation(receiver, receiverServer)
		if err != nil{
			return err
		}
		convo = newConvo
	}
	cli.messageBuilder.
		SetMsgString(message).
		SetMsgContentBuilder(convo.MessageBuilder()).
		SetReceiverEncrypter(convo.Encrypter()).
		SetReceiver(receiver).SetReceiverServer(receiverServer).
		SetPath(cli.currentPath)

	msg, err := cli.messageBuilder.Build()
	if err != nil {
		return err
	}

	err = cli.Send(msg)
	if err != nil {
		return err
	}
	return nil
}

func (cli *Client)ReceiveMessage(content []byte, receiver string, receiverServer string)error{
	name := receiverServer + receiver
	cli.convosMutex.Lock()
	convo, ok := cli.conversations[name]
	cli.convosMutex.Unlock()
	if !ok{
		newConvo, err := cli.CreateConversation(receiver, receiverServer)
		if err != nil{
			return err
		}
		convo = newConvo
	}
	convo.Receive(content)
	return nil
}


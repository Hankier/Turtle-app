package client

import (
	"session"
	"serverList"
	"log"
	"net"
	"messageHandler"
	"cryptographer"
	"errors"
	"message"
	"textReceiver"
	"conversation"
	"messageBuilder"
	"sync"
)

type Client struct{
	sess           *session.Session
	srvList        *serverList.ServerList
	nodeCrypto     cryptographer.Cryptographer
	currentPath    []string
	convosMutex	   sync.Mutex
	conversations  map[string]*conversation.Conversation
	messageBuilder *messageBuilder.MessageBuilder
	textReceiver   textReceiver.TextReceiver
	myName         string
}

func NewClient()(*Client){
	cli := new(Client)

	cli.srvList = serverList.NewServerList()
	cli.nodeCrypto = cryptographer.NewNodeCrypto()
	cli.textReceiver = &textReceiver.TextReceiverImpl{}

	return cli
}

func (cli *Client)CreateSession(name string, socket net.Conn){
	if cli.sess != nil{
		cli.RemoveSession()
	}
	msgHandler := messageHandler.NewMessageHandlerImpl(cli, cli.nodeCrypto)

	sess := session.NewSession(socket, name, msgHandler, cli)

	go sess.Start()
	cli.sess = sess
	//TODO thread safe
}

func (cli *Client)RemoveSession(){
	cli.sess.DeleteSession()
	cli.sess = nil
}

func (cli *Client)Send(content []byte, receiver string, receiverServer string)error{
	msg, err := cli.messageBuilder.SetMsgContent(content).SetReceiver(receiver).SetReceiverServer(receiverServer).Build()
	if err != nil{
		return err
	}
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

func (cli *Client)ChooseNewPath(length int)[]string{
	cli.currentPath = cli.srvList.GetRandomPath(length)
	return cli.currentPath
}

func (cli *Client)ConnectToServer(name string)error{
	socket, err := net.Dial("tcp", cli.srvList.GetServerIpPort(name))
	if err != nil {
		return err
	}
	socket.Write([]byte(cli.myName))
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
		convo = conversation.NewConversation(cli.textReceiver, receiver, receiverServer)
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
	convoMessage := convo.BuildMessage(cli.myName, cli.sess.GetName(), message)

	err := cli.Send(convoMessage.ToBytes(), receiver, receiverServer)
	if err != nil {
		return err
	}
	return nil
}

func (cli *Client)ReceiveMessage(content []byte, receiver string, receiverServer string){
	name := receiverServer + receiver
	cli.convosMutex.Lock()
	convo, ok := cli.conversations[name]
	if !ok{
		convo = conversation.NewConversation(cli.textReceiver, receiver, receiverServer)
		cli.conversations[name] = convo
	}
	cli.convosMutex.Unlock()
	convo.Receive(content)
}


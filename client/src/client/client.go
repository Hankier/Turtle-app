package client

import (
	"srvlist"
	"log"
	"net"
	"textReceiver"
	"sync"
	"msgs/builder"
	"sessions"
	"convos"
)

type Client struct{
	myName         string

	srvList        *srvlist.ServerList
	sessionsContr  *sessions.Controller
	convosContr	   *convos.Controller
	currentPath    []string
	convosMutex    sync.Mutex
	msgsBuilder    *builder.Builder
	textReceiver   textReceiver.TextReceiver
}

func New(name string)(*Client){
	cli := new(Client)

	cli.myName = name

	cli.srvList = srvlist.New()
	cli.textReceiver = &textReceiver.TextReceiverImpl{}
	cli.convosContr = convos.New(cli.textReceiver)
	cli.sessionsContr = sessions.New(cli.convosContr)
	cli.msgsBuilder = builder.New(cli, cli.srvList)

	return cli
}

func (cli *Client)Start(){
	cli.cmdListener.Listen()
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

	srvPort, err := cli.srvList.GetServerIpPort(name)
	if err != nil {	return err	}

	socket, err := net.Dial("tcp", srvPort)
	if err != nil {	return err	}

	socket.Write([]byte(cli.myName))
	cli.msgsBuilder.SetMyCurrentServer(name)
	cli.CreateSession(name, socket)
	log.Print("Succesfully connected to " + name)
	return nil
}

func (cli *Client)GetServerList()[]string{
	return cli.srvList.GetServerList()
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
	cli.msgsBuilder.
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

func (cli *Client)OnReceive(from string, content []byte){
	cli.convosContr.OnReceive(from, content)
}


func (cli *Client)GetName()string{
	return cli.myName
}

func (cli *Client)GetCurrentServer()(string, error){
	sessionsNames := cli.sessionsContr.GetActiveSessions()
	if len(sessionsNames) < 1{
		return "", errors.New("no active session")
	}
	return sessionsNames[0], nil
}
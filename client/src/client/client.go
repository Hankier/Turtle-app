package client

import (
	"srvlist"
	"log"
	"net"
	"textReceiver"
	"msgs/builder"
	"sessions"
	"convos"
	"errors"
	"cmdsListener"
)

type Client struct{
	myName         string
	srvList        *srvlist.ServerList
	sessionsContr  *sessions.Controller
	convosContr	   *convos.Controller
	currentPath    []string
	msgsBuilder    *builder.Builder
	textReceiver   textReceiver.TextReceiver
	commandsListener *cmdsListener.Listener
}

func New(name string)(*Client){
	cli := new(Client)

	cli.myName = name

	cli.srvList = srvlist.New()
	cli.textReceiver = &textReceiver.TextReceiverImpl{}
	cli.convosContr = convos.New(cli.textReceiver)
	cli.sessionsContr = sessions.New(cli.convosContr)
	cli.msgsBuilder = builder.New(cli.srvList, cli.convosContr, cli)
	cli.commandsListener = cmdsListener.New(cli, cli.textReceiver)

	return cli
}

func (cli *Client)Start(){
	cli.commandsListener.Listen()
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
	activeSessions := cli.sessionsContr.GetActiveSessions()
	for i := 0; i < len(activeSessions); i++{
		cli.sessionsContr.RemoveSession(activeSessions[i])
	}
	cli.sessionsContr.CreateSession(name, socket)
	log.Print("Succesfully connected to " + name)
	return nil
}

func (cli *Client)GetServerList()[]string{
	return cli.srvList.GetServerList()
}

func (cli *Client)SendTo(receiverServer string, receiver string, command string)error{

	cli.msgsBuilder.SetCommand(command).
		SetReceiver(receiver).SetReceiverServer(receiverServer).
		SetPath(cli.currentPath)

	message, err := cli.msgsBuilder.Build()
	if err != nil {
		return err
	}

	currentServer, err := cli.GetCurrentServer()
	if err != nil {
		return err
	}

	err = cli.sessionsContr.Send(currentServer, message.ToBytes())
	if err != nil {
		return err
	}
	return nil
}

func (cli *Client)OnReceive(from string, content []byte){
	cli.convosContr.OnReceive(from, content)
}

func (cli *Client)CreateConversation(receiverServer string, receiver string) (err error){
	return cli.convosContr.CreateConversation(receiverServer, receiver)
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
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
	"strconv"
	"strings"
	"reflect"
)

/*
Client class implementing UserInterface, (sessions)Receiver, (conversation)Receiver and CredentialHolder interfaces
Main class of the program
*/
type Client struct{
	myName         string					//name of self
	srvList        *srvlist.ServerList		//serverList object for operations on servers
	sessionsContr  *sessions.Controller		//controller for operations on sessions
	convosContr	   *convos.Controller		//controller for operations on conversations
	currentPath    []string					//slice of server names representing current path to send messages by
	msgsBuilder    *builder.Builder			//builder object for creating ready-to-send messages to send
	textReceiver   textReceiver.TextReceiver//text receiver to forward to controllers
	commandsListener *cmdsListener.Listener	//listener handling given commands
}

//New is a constructor
//Creates Client with a given name
//Initializes empty textReceiver Object,
//serverList created via its constructor
//conversationsController crated via its constructor, with created textReceiver and self (as CredentialsHolder) as parameters
//sessionsController created via its constructor, with created convosContr as parameter
//messageBuilder created via its constructor, with serverList, convosContr and self (as CredentialsHolder) as parameters
//commandsListener created via its constructor, with self(as UserInterface) and textReceiver as parameters
func New(name string)(*Client){
	cli := new(Client)

	cli.myName = name

	cli.srvList = srvlist.New()
	//TODO remove debug data
	cli.srvList.DebugGetServers()

	cli.textReceiver = &textReceiver.TextReceiverImpl{}
	cli.convosContr = convos.New(cli.textReceiver, cli)
	cli.sessionsContr = sessions.New(cli.convosContr)
	cli.msgsBuilder = builder.New(cli.srvList, cli.convosContr, cli)
	cli.commandsListener = cmdsListener.New(cli, cli.textReceiver)

	return cli
}

//Starts listening for commands
func (cli *Client)Start(){
	cli.commandsListener.Listen()
}

//GetCurrentPath eturns a slice of server names as strings representing consecutive nodes of path in reverse order
func (cli *Client)GetCurrentPath() []string{
	return cli.currentPath
}

//ChooseNewPath generates a new random path and assigns it as a current path
//Returns nil and error if serverList object encountered a problem generating a path
//Returns generated path and nil if all went well
func (cli *Client)ChooseNewPath(length int)([]string, error){
	var err error
	cli.currentPath, err = cli.srvList.GetRandomPath(length)

	if err != nil{
		return nil, err
	}

	return cli.currentPath, nil
}

//ConnectToServer tries to connect to server of a given name and writes self's name to it
//Removes all active sessions
//Furthermore it creates a new session with the server it newly connected to
//Returns error if there is no server named as required or there are problems with connection
//Returns nil if all went well
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

//GetServerList returns a slice of all known server names as strings
func (cli *Client)GetServerList()[]string{
	return cli.srvList.GetServerList()
}

func (cli *Client)GetServerDetails(name string)[]string{
	details := make([]string, 1)
	details[0], _ = cli.srvList.GetServerIpPort(name)
	encrypter, _ := cli.srvList.GetEncrypter(name)
	rsaKey := encrypter.GetPublicKeyRSA()
	elg := encrypter.GetPublicKeyElGamal()
	if rsaKey != nil{
		details = append(details, "rsa: " + strconv.Itoa(rsaKey.E))
	}
	if elg != nil{
		details = append(details, "elgamal: " + elg.Y.String())
	}
	return details
}

//SendTo tries to send a message specified in a command to a receiver which should be connected to given receiverServer
//Returns error in case:
// -builder object encountered a problem creating a ready-to-send message from given parameters
// -client is not connected to any server
// -sessionsController encountered a problem sending message
//Return nil if all went well
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

//OnReceive passes received content from another server to convosController method OnReceive
func (cli *Client)OnReceive(from string, content []byte){
	cli.convosContr.OnReceive(from, content)
}

//CreateConversation passes given parameters to convosController CreateConversation method
//Returns error accordingly to that function
func (cli *Client)CreateConversation(receiverServer string, receiver string) (err error){
	return cli.convosContr.CreateConversation(receiverServer, receiver)
}

//GetName returns client's name
func (cli *Client)GetName()string{
	return cli.myName
}

//GetCurrentServer returns name of a server client is currently connected to
//Returns empty string and an error if client is not connected to any server
func (cli *Client)GetCurrentServer()(string, error){
	sessionsNames := cli.sessionsContr.GetActiveSessions()
	if len(sessionsNames) < 1{
		return "", errors.New(reflect.TypeOf(cli).String() + ": no active session")
	}
	return sessionsNames[0], nil
}
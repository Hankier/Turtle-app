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
)

type Client struct{
	sess *session.Session
	srvList	*serverList.ServerList
	clientCrypto *cryptographer.ClientCrypto
	myName string
}

func NewClient()(*Client){
	cli := new(Client)

	cli.srvList = serverList.NewServerList()

	return cli
}


func (cli *Client)SendTo(msg *message.Message)error{

	if cli.sess != nil{
		cli.sess.Send(msg)
		return nil
	}else{
		log.Println("Not connected to any server\n");
		return errors.New("NOT CONNECTED")
	}
}

func (cli *Client)SendInstantTo(msg *message.Message)error{
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

func (cli *Client)ConnectToServer(name string)bool{
	server := cli.serverList[name]
	socket, err := net.Dial("tcp", server.Ip_port)
	if err != nil {
		log.Print("Error connecting to server " + name + " " + err.Error())
		return false
	}
	socket.Write([]byte(cli.myName))
	cli.CreateSession(name, socket)
	log.Print("Succesfully connected to " + name)
	return true
}

func (cli *Client)CreateSession(name string, socket net.Conn){
	if cli.sess != nil{
		cli.RemoveSession()
	}
	msgHandler := messageHandler.NewMessageHandlerImpl(cli, cli.clientCrypto)

	sess := session.NewSession(socket, name, msgHandler, cli)

	go sess.Start()
	cli.sess = sess
	//TODO thread safe
}

func (cli *Client)RemoveSession(){
	cli.sess.DeleteSession()
	cli.sess = nil
}


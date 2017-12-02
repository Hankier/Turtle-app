package client

import (
	"session"
	"serverEntry"
	"log"
	"net"
	"messageHandler"
	"cryptographer"
	"crypto/rand"
	"math/big"
)

type Client struct{
	sess *session.Session
	serverList map[string]*serverEntry.ServerEntry
	clientCrypto *cryptographer.ClientCrypto
	myName string
}

func NewClient()(*Client){
	cli := new(Client)

	cli.serverList = make(map[string]*serverEntry.ServerEntry)

	pk := make([]byte, 256)

	cli.serverList["00000000"] = serverEntry.NewServerEntry("00000000", "127.0.0.1:8081", pk)
	cli.serverList["00000001"] = serverEntry.NewServerEntry("00000001", "127.0.0.1:8083", pk)
	cli.serverList["00000002"] = serverEntry.NewServerEntry("00000002", "127.0.0.1:8085", pk)

	return cli
}


func (cli *Client)SendTo(bytes []byte)error{
	if cli.sess != nil{
		cli.sess.Send(bytes)
		return nil
	}else{
		log.Println("Not connected to any server\n");
		return new(error)
	}
}

func (cli *Client)SendInstantTo(bytes []byte)error{
	if cli.sess != nil{
		cli.sess.SendInstant(bytes)
		return nil
	}else{
		log.Println("Not connected to any server\n");
		return new(error)
	}
}

func (cli *Client)UnlockSending(){
	cli.sess.UnlockSending()
}

func (cli *Client)connectToServer(name string)bool{
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
	msgHandler := messageHandler.NewMessageHandlerImpl(cli, cli.clientCrypto)

	sess := session.NewSession(socket, name, msgHandler, cli)

	go sess.Start()
	cli.sess = sess
	//TODO thread safe
}

func (cli *Client)RemoveSession(){
	cli.sess.DeleteSession()
}

func (cli *Client)GetRandomPath(length int)[]*serverEntry.ServerEntry{
	path := make([]*serverEntry.ServerEntry, length)

	keys := make([]string, 0, len(cli.serverList))
	for k := range cli.serverList {
		keys = append(keys, k)
	}


	for i,_ := range path{
		rnd, _ := rand.Int(rand.Reader, big.NewInt(int64(len(keys))))
		key := keys[rnd.Int64()]
		path[i] = cli.serverList[key]
	}

	return path
}
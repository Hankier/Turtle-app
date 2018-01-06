package server

import (
	"sync"
	"log"
	"net"
	"srvlist"
	"sessions"
	"errors"
	"server/listener"
)

type Server struct{
	myName string
	serverList *srvlist.ServerList
	sessionsContr *sessions.Controller
	clientListener *listener.Listener
	serverListener *listener.Listener
	wg sync.WaitGroup
}

func NewServer(name string)(*Server){
	srv := new(Server)

	srv.myName = name
	//TODO Downloading server list from DA
	srv.serverList = srvlist.New()
	//TODO remove debug data
	srv.serverList.DebugGetServers()

	srv.sessionsContr = sessions.New(srv, srv)
	srv.wg.Add(2)
	return srv
}


func (srv *Server)Start(clientPort, serverPort string)error{
	var err error
	srv.clientListener, err = listener.New(clientPort, srv.sessionsContr)
	if err != nil {
		log.Fatal(err)
	}
	srv.serverListener, err = listener.New(serverPort, srv.sessionsContr)
	if err != nil {
		log.Fatal(err)
	}
	if err == nil{
		go srv.clientListener.Loop(srv.wg)
		go srv.serverListener.Loop(srv.wg)
		log.Print("Server started")
		srv.wg.Wait()
	}

	return err
}


func (srv *Server)ConnectToServer(name string)error{

	if ipport, ok := srv.serverList.GetServerIpPort(name); ok == nil{
		socket, err := srv.dialAndSendName(name, ipport)
		if err != nil {
			return err
		}
		srv.sessionsContr.CreateSession(name, socket)
		log.Print("Succesfully connected to " + name)
	} else {
		return errors.New("no server on list " + name)
	}
	return nil
}

func (srv *Server)dialAndSendName(name, ipport string)(net.Conn, error){
	socket, err := net.Dial("tcp", ipport)
	if err != nil {
		log.Print("Error connecting to server " + name + " " + err.Error())
		return nil, err
	}
	socket.Write([]byte(srv.myName))
	return socket, nil
}

func (srv *Server)checkIfNameIsServer(name string)bool{
	for _, server := range srv.serverList.GetServerList(){
		if server == name{
			return true
		}
	}
	return false
}

func (srv *Server)GetName()string{
	return srv.myName
}
package server

import (
	"connectionListener"
	"sync"
	"decrypter"
	"log"
	"net"
	"serverEntry"
	"sessionHandler"
	"session"
	"messageHandler"
)

type Server struct{
	sessions map[string]*session.Session
	clientListener *connectionListener.ConnectionListener
	serverListener *connectionListener.ConnectionListener
	serverList map[string]serverEntry.ServerEntry
	serverCrypto *decrypter.ServerCrypto
	wg sync.WaitGroup
	myName string
}

func NewServer(name string)(*Server){
	srv := new(Server)

	srv.sessions = make(map[string]*session.Session)

	srv.myName = name

	//TODO Downloading server list from hardcoded DA

	srv.wg.Add(2)
	srv.serverCrypto = decrypter.NewServerCrypto();
	return srv
}

func checkIfNameIsServer(name string)bool{
	//TODO
	return true;
}

func (srv *Server)SendTo(name string, bytes []byte){
	if session, ok := srv.sessions[name]; ok {
		session.Send(bytes)
	}else{
		if checkIfNameIsServer(name) {
			if srv.connectToServer(name){
				srv.SendTo(name, bytes)
			}
		}
	}
}

func (srv *Server)UnlockSending(name string){}


func (srv *Server)Start(clientPort, serverPort string)error{
	var err error
	srv.clientListener, err = connectionListener.NewConnectionListener(clientPort, srv)
	if err != nil {
		log.Fatal(err)
	}
	srv.serverListener, err = connectionListener.NewConnectionListener(serverPort, srv)
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

func (srv *Server)connectToServer(name string)bool{
	server := srv.serverList[name]
	socket, err := net.Dial("tcp", server.Ip_port)
	if err != nil {
		log.Fatalln("Error connecting to server" + name)
		return false
	}
	socket.Write([]byte(srv.myName))
	srv.CreateSession(name, &socket)
	return true
}

func (srv *Server)CreateSession(name string, socket *net.Conn){
	msgHandler := messageHandler.NewMessageHandlerImpl(srv, srv.serverCrypto)
	sess := session.NewSession(socket, name, msgHandler)

	sess.Start()
	srv.sessions[name] = sess
}

func (srv *Server)RemoveSession(name string){
	//TODO
}
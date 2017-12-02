package server

import (
	"connectionListener"
	"sync"
	"decrypter"
	"log"
	"net"
	"serverEntry"
	"session"
	"messageHandler"
)

type Server struct{
	sessions map[string]*session.Session
	clientListener *connectionListener.ConnectionListener
	serverListener *connectionListener.ConnectionListener
	serverList map[string]*serverEntry.ServerEntry
	serverCrypto *decrypter.ServerCrypto
	wg sync.WaitGroup
	myName string
}

func NewServer(name string)(*Server){
	srv := new(Server)

	srv.sessions = make(map[string]*session.Session)

	srv.myName = name

	//TODO Downloading server list from DA

	srv.serverList = make(map[string]*serverEntry.ServerEntry)

	pk := make([]byte, 256)

	srv.serverList["00000000"] = serverEntry.NewServerEntry("00000000", "127.0.0.1:8081", pk)
	srv.serverList["00000001"] = serverEntry.NewServerEntry("00000001", "127.0.0.1:8083", pk)
	srv.serverList["00000002"] = serverEntry.NewServerEntry("00000002", "127.0.0.1:8085", pk)
	srv.wg.Add(2)
	srv.serverCrypto = decrypter.NewServerCrypto();
	return srv
}

func checkIfNameIsServer(name string)bool{
	if name[0] == '0'{
		return true;
	}else {
		return false
	}
}

func (srv *Server)SendTo(name string, bytes []byte){
	if sess, ok := srv.sessions[name]; ok {
		sess.Send(bytes)
	}else{
		if checkIfNameIsServer(name) {
			if srv.connectToServer(name){
				srv.SendTo(name, bytes)
			}
		}
	}
}

func (srv *Server)SendInstantTo(name string, bytes []byte){
	if sess, ok := srv.sessions[name]; ok {
		sess.SendInstant(bytes)
	}else{
		if checkIfNameIsServer(name) {
			if srv.connectToServer(name){
				srv.SendTo(name, bytes)
			}
		}
	}
}

func (srv *Server)UnlockSending(name string){
	srv.sessions[name].UnlockSending()
}


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
		log.Print("Error connecting to server " + name + " " + err.Error())
		return false
	}
	socket.Write([]byte(srv.myName))
	srv.CreateSession(name, socket)
	log.Print("Succesfully connected to " + name)
	return true
}

func (srv *Server)CreateSession(name string, socket net.Conn){
	msgHandler := messageHandler.NewMessageHandlerImpl(srv, srv.serverCrypto)

	sess := session.NewSession(socket, name, msgHandler, srv)

	go sess.Start()
	srv.sessions[name] = sess
	//TODO thread safe
}

func (srv *Server)RemoveSession(name string){
	srv.sessions[name].DeleteSession()
	delete(srv.sessions, name)
	//TODO thread safe
}
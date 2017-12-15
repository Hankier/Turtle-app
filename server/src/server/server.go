package server

import (
	"connectionListener"
	"sync"
	"cryptographer"
	"log"
	"net"
	"srvlist"
	"sessions"
	"messageHandler"
	"message"
)

type Server struct{
	sessions struct{
		sync.Mutex
		data map[string]*sessions.Session
	}
	clientListener *connectionListener.ConnectionListener
	serverListener *connectionListener.ConnectionListener
	serverList *srvlist.ServerList
	serverCrypto *cryptographer.NodeCrypto
	wg sync.WaitGroup
	myName string
}

func NewServer(name string)(*Server){
	srv := new(Server)

	srv.sessions.data = make(map[string]*sessions.Session)

	srv.myName = name

	//TODO Downloading server list from DA

	srv.serverList = srvlist.New()

	srv.wg.Add(2)
	srv.serverCrypto = cryptographer.New();
	return srv
}

func (s *Server)checkIfNameIsServer(name string)bool{
	for _, server := range s.serverList.GetServerList(){
		if server == name{
			return true
		}
	}
	return false
}

func (srv *Server)SendTo(name string, msg *message.Message){
	srv.sessions.Lock()
	sess, ok := srv.sessions.data[name];
	srv.sessions.Unlock()

	if ok {
		sess.Send(msg)
	}else{
		if srv.checkIfNameIsServer(name) {
			if srv.connectToServer(name){
				srv.SendTo(name, msg)
			}
		}
	}
}

func (srv *Server)SendInstantTo(name string, msg *message.Message){
	srv.sessions.Lock()
	sess, ok := srv.sessions.data[name]
	srv.sessions.Unlock()

	if ok {
		sess.SendInstant(msg)
	}else{
		if srv.checkIfNameIsServer(name) {
			if srv.connectToServer(name){
				srv.SendTo(name, msg)
			}
		}
	}
}

func (srv *Server)UnlockSending(name string){
	srv.sessions.Lock()
	sess := srv.sessions.data[name]
	srv.sessions.Unlock()

	sess.UnlockSending()
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

	if ipport, ok := srv.serverList.GetServerIpPort(name); ok == nil{
		socket, err := srv.dialAndSendName(name, ipport)
		if err != nil {return false}
		srv.CreateSession(name, socket)
		log.Print("Succesfully connected to " + name)
		return true
	}
	log.Print("No server on list " + name)
	return false
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

func (srv *Server)CreateSession(name string, socket net.Conn){
	msgHandler := messageHandler.NewMessageHandlerImpl(srv, srv.serverCrypto)

	sess := sessions.New(socket, name, msgHandler, srv)

	go sess.Start()

	srv.sessions.Lock()
	srv.sessions.data[name] = sess
	srv.sessions.Unlock()
}

func (srv *Server)RemoveSession(name string){
	srv.sessions.Lock()
	srv.sessions.data[name].DeleteSession()
	delete(srv.sessions.data, name)
	srv.sessions.Unlock()
}
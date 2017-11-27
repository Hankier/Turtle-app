package server

import (
	"connectionListener"
	"sync"
	"serverCrypto"
	"log"
	"fmt"
	"net"
	"strings"
)

type Server struct{
	sessions map[[256]byte]string
	clientListener *connectionListener.ConnectionListener
	serverListener *connectionListener.ConnectionListener
	serverList map[string][256]byte
	serverCrypto *serverCrypto.ServerCrypto
	wg sync.WaitGroup
}

func MakeServer()(*Server){
	srv := new(Server)

	srv.wg.Add(2)
	srv.serverCrypto = serverCrypto.NewServerCrypto();
	return srv
}

func (srv *Server)sendTo(ip string,  message []byte) {

	fmt.Println(ip)



	connection, err := net.Dial("tcp", ip )
	if err != nil {
		log.Fatal(err)
	}
	connection.Write(message)
}

func getPort(ip string)string{
	return strings.Split(ip, ":")[1]
}

func (srv *Server)Start()error{
	var err error
	srv.clientListener, err = connectionListener.NewConnectionListener("4000", nil)//TODO handler
	if err != nil {
		log.Fatal(err)
	}
	srv.serverListener, err = connectionListener.NewConnectionListener("4001", nil)//TODO handler
	if err != nil {
		log.Fatal(err)
	}
	if err == nil{
		go srv.clientListener.Loop(srv.wg)
		go srv.serverListener.Loop(srv.wg)
	}
	srv.wg.Wait()

	return err
}
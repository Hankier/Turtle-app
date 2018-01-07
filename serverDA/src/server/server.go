package server

import (
	_ "sync"
	"log"
	"net"
	"serverlist"
	_ "errors"
	"encoding/json"

)

type Server struct{
	myName string
	serverList *serverlist.ServerList
	//clientListener *listener.Listener
	//serverListener *listener.Listener
	//wg sync.WaitGroup
}


type msgMessage struct{
	Type    string
	Content string
	Content2 string
}

func NewServer(name string)(*Server){
	srv := new(Server)

	srv.myName = name
	srv.serverList = serverlist.NewList()
	return srv
}



func (srv *Server)Start(port string){
	log.Println("Port: ", port)
	ln, err := net.Listen("tcp", ":"+port)
    if err != nil {
        log.Fatal(err)
    }

    log.Println("Server up and listening on port ", port)


    for {
        conn, err := ln.Accept()
        if err != nil {
            log.Println(err)
            continue
        }
        go handleConnection(conn, srv.serverList)
	}
}

func handleConnection(c net.Conn, list *serverlist.ServerList) {

    log.Printf("Client %v connected.", c.RemoteAddr())

	//to read json from stream
	d := json.NewDecoder(c)

    var msg msgMessage

    err := d.Decode(&msg)
    message_type := msg.Type
    message_content := msg.Content
	message_content2 := msg.Content2
	log.Println("Errors: ",err)

	if message_type == "ADD" {
		list.AddServerToList(message_content, message_content2)
	} else {
		log.Println("Not know type: ",message_type)
	}

	aa, _ := list.GetServerIpPort(list.GetServerList()[0])
	log.Printf("Serverlist added in handler: %s", aa)

    c.Close()

    //messageBuffer := make([]byte, 2048)
    //for {
    //  n, err := c.Read(messageBuffer)
    //  log.Println(err)
    //  if err != nil {
    //      c.Close()
    //      break
    //  }
    //  msg := string(messageBuffer)
      //sendTo(ip, msg)

      //log.Printf("%v", )
      //log.Printf("Data: %v , %v", n, messageBuffer[0:n])

    log.Printf("Connection from %v closed.", c.RemoteAddr())
 }


// func (srv *Server)Start(clientPort, serverPort string)error{
// 	var err error
// 	srv.clientListener, err = listener.NewConnectionListener(clientPort)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	srv.serverListener, err = listener.NewConnectionListener(serverPort)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	if err == nil{
// 		go srv.clientListener.Loop(srv.wg)
// 		go srv.serverListener.Loop(srv.wg)
// 		log.Print("Server started")
// 		srv.wg.Wait()
// 	}
//
// 	return err
// }
//




func (srv *Server)GetName()string{
	return srv.myName
}

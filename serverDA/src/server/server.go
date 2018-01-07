package server

import (
	_ "sync"
	"log"
	"net"
	"serverlist"
	_ "errors"
	"encoding/json"
	"io/ioutil"

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

type response struct{
    Status string
    Name string
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
	srv.serverList.AddServerToList("1.2.3.4:123456", `-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEA0dkp8n9H3c7/OlaYKpi8Fu7SWsqF3eeRw3C8gpLPkMhTDFRm
/ZpytL/ImYAh1mQhvvH5kHPHc7unARSDlaa/OzDcMqwSlXmzIWw1k81a3HQj+1wB
anbIftyo/Pc1WwKQ35ppCtwJWytCipeqD3+ACEbhVmb+4TTHX3/MO5E1mWFNISJW
VNTMxIibqBPG2cDjRAu4kGNaVM1pd1/AZrid6LtvmUaXajLhodywtopBGUana4qR
5ou9OPYc3gELMkMQCM4AUqWM04iJ0+620rmoeFHRu3L4miekU+bwZiDiNXG4wXDp
HgmbStRVUhLHntYGs+4bZBfghW5WZqDxNsIMPwIDAQABAoIBAQCXftDyuYLXlgXa
RwPJ1MQNRlLkqsrj/bbUwsHE/loNKyIRh6lmsqbW6JHYh5FmJpnaMPS7nWpDmhii
Bf5M/rmV8Ns3VdSAxwBUQ7uWPa2387y6TZzUEHcEZyc0oP+K+Zo/Y0ksRtgWUm/S
gFWMpL54uzsY1nhxe1noDuoRou5wEGkmatlXr0FvME4OmwOe1S24bQg97A9dqngh
PELJ+pdQKcSqWAgW8Y3PMxx88DI08DOF1IxyV8VvTrWhbzJvTZnpQqxy3pDFr5TO
S287PicRAeCyDM5Rxi6tzJaLuh19RT4YtKZ1qgJRQUEOtM6Pf0rT7/YAjiFdpoT0
MNNboGCBAoGBAOdqwI+KEvPf5RXaTeiZcymS9fwGCyObDfOvtrOhKAc+cv5wNwJG
g6LsK+i/5EWi9INEd8O1WcYVaqjSG6r36zxjcWOKrRNU1gyvE8ONg+DuQ34nZz0V
rjgnzSIzK4OlOdPpfU/wKuVipC0quY2HGODxrV7KtaU4++mr4bIefEPBAoGBAOgj
3RK6t2jCABmGG1iruY2/YEhT+blOJ9A6YKMptfWMwuewL5A11vH1j00JKm6N1twQ
kHKSdUVurZArwArQ7LxLNa2VhTzjMIviMtnTQc+0CdkUKNGeCvx2JVQbzr8Av2eg
r/Dt7weSBLerWFq2cAwvrRZ665sx6cRU3QTckk//AoGAe/kXgY4hix6F1kgl9pbG
OB5vwvzl2MRHHCYlBWQvUnolFqO9BG4MNSq6dyzduGSNAwmZ83Fiz5hHlHtCsTux
fJ91bjMrdzC6nv7n4pocbVKXO60WRIYp2BGSdmDdTeAk856hMELkaBCJDV1XHDek
n1U5YI/N8d5uLgeTmF12isECgYEAq1F8X8woe0lhJXURTXk+cVvhRL+ktpr1Svkq
RIAN52/Aj5g5IeZ6AQtGfIXdKMXI4ZPf5o4rudgagyGmktTpQXUH4llMgUjxlOqU
uKjuEsk901TLYxeN6A+RMOdsxw1YNLQj5FzUYPPkQ2BSzm+BdZzh0otYwaouaVRv
4Jyf5iUCgYAFAzx2WWGhDGnyMweiEfKmwoKDaOAShDadvFmwy5irTKiceflQRhHJ
obORGt2A/HtVMpiG+4fjHnAMaKWahaTJ5FKgaRHWahYjVsPRRJwrqbLILEo9Q90N
CuHDJhgWQwJl6/9NXxbFTrCcBRTd5SHvptaUh+ZponwS/ltLdi6bPQ==
-----END RSA PRIVATE KEY-----`)
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
	//test := string(c.RemoteAddr())
	//log.Printf("Test: %s", test)

	//to read json from stream
	d := json.NewDecoder(c)

    var msg msgMessage

    err := d.Decode(&msg)
    message_type := msg.Type
    message_content := msg.Content
	message_content2 := msg.Content2
	log.Println("Errors: ",err)

	if message_type == "ADD" {
		name := list.AddServerToList(message_content, message_content2)
		enc := json.NewEncoder(c)
		res := response{
		Status:   "OK",
		Name:	name,
	}
    enc.Encode(res)
	} else if message_type == "GET" {
		//enc := json.NewEncoder(c)
		res := list.GetServerListJSON()
    c.Write(res)
	//data_j, _ := json.Marshal(res)
	ioutil.WriteFile("test.json",res,0644)

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

package listener

import (
	"net"
	"log"
	"sync"
	"io"
	"sessions/handler"
)


type Listener struct{
	socket          net.Listener
	sessionsHandler handler.Handler
}

func New(port string, sessionsHandler handler.Handler) (*Listener, error) {
	cln := new(Listener)
	port = ":"+port
	var err error
	cln.socket, err = net.Listen("tcp", port)
	cln.sessionsHandler = sessionsHandler
	return cln, err
}


func (cln *Listener)handleConnection(c net.Conn) {
	log.Printf("Client %v connected to port %d", c.RemoteAddr(), cln.socket.Addr().(*net.TCPAddr).Port)

	nameBytes := make([]byte, 8)
	io.ReadFull(c, nameBytes)

	name := string(nameBytes)

	cln.sessionsHandler.CreateSession(name, c)
}

func (cln *Listener)Loop(wg sync.WaitGroup) error{
	defer wg.Done()
	for{
		for {
			conn, err := cln.socket.Accept()
			if err != nil {
				return err
			}
			go cln.handleConnection(conn)
		}
	}
	return nil
}
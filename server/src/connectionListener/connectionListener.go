package connectionListener

import (
	"net"
	"log"
	"sync"
	"io"
	"sessions/handler"
)


type ConnectionListener struct{
	socket          net.Listener
	sessionsHandler handler.Handler
}

func NewConnectionListener(port string, sessionsHandler handler.Handler) (*ConnectionListener, error) {
	cln := new(ConnectionListener)
	port = ":"+port
	var err error
	cln.socket, err = net.Listen("tcp", port)
	cln.sessionsHandler = sessionsHandler
	return cln, err
}


func (cln *ConnectionListener)handleConnection(c net.Conn) {
	log.Printf("Client %v connected to port %d", c.RemoteAddr(), cln.socket.Addr().(*net.TCPAddr).Port)

	nameBytes := make([]byte, 8)
	io.ReadFull(c, nameBytes)

	name := string(nameBytes)

	cln.sessionsHandler.CreateSession(name, c)
}

func (cln *ConnectionListener)Loop(wg sync.WaitGroup) error{
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
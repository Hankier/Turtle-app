package connectionListener

import (
	"net"
	"fmt"
	"log"
	"encoding/binary"
	"regexp"
	"sync"
)

type SessionHandler interface{
	createSession(conn net.IPConn)
	removeSession(ip net.IP)
}


type ConnectionListener struct{
	socket net.Listener
	sessionsHandler SessionHandler
	wg sync.WaitGroup
}

func NewConnectionListener(port string, handler SessionHandler) (*ConnectionListener, error) {
	cln := new(ConnectionListener)
	port = ":"+port
	var err error
	cln.socket, err = net.Listen("tcp", port)
	cln.sessionsHandler = handler
	cln.wg.Add(1)
	return cln, err
}


func (cln *ConnectionListener)handleConnection(c net.Conn) {

	log.Printf("Client %v connected.", c.RemoteAddr())

	messageSizeBuffer := make([]byte, 4)

	n, err := c.Read(messageSizeBuffer)
	if err != nil || n == 0 {
		c.Close()
	}
	messageSize := binary.LittleEndian.Uint32(messageSizeBuffer[0:])

	log.Printf("Data1: %v ", messageSize)
	log.Printf("Data: %v , %v", n, messageSizeBuffer[0:n])

	log.Printf(string(messageSizeBuffer[0:n]));

	messageBuffer := make([]byte, messageSize)
	for {
		n, err := c.Read(messageBuffer)
		fmt.Println(err)
		if err != nil {
			c.Close()
			break
		}
		ip, msg := getIp(string(messageBuffer))
		//sendTo(ip, msg)

		log.Printf(ip)
		log.Printf("%v", msg[0:])
		log.Printf("Data: %v , %v", n, messageBuffer[0:n])

	}
	log.Printf("Connection from %v closed.", c.RemoteAddr())
}

func (cln *ConnectionListener)Loop() error{
	defer cln.wg.Done()
	for{
		fmt.Println("Server up and listening on port 12345")

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

func (cln *ConnectionListener)Join(){
	cln.wg.Wait()
}

func getIp(message string) (ipAddress string, newMessage []byte)  {
	re := regexp.MustCompile("(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5]):[0-9]+ ")
	ipAddress = re.FindString(message)
	fmt.Println(ipAddress)
	reg := regexp.MustCompile(ipAddress)
	message = reg.ReplaceAllString(message, "")
	newMessage = []byte(message)

	return ipAddress, newMessage
}

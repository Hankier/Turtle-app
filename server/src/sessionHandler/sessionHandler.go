package sessionHandler


import "net"

type SessionHandler interface{
	CreateSession(name string, socket net.Conn)
	RemoveSession(name string)
}

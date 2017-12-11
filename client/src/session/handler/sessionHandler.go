package handler


import "net"

type Handler interface{
	CreateSession(name string, socket net.Conn)
	RemoveSession()
}

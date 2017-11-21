package sessionHandlerInterface


import "net"

type SessionHandler interface{
	createSession(conn net.IPConn)
	removeSession(ip net.IP)
}

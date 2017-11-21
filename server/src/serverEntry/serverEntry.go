package serverEntry

type ServerEntry struct{
	name string
	ip_port string
	publicKey [256]byte
}

func NewServerEntry(name string, ip_port string, publicKey [256]byte)(*ServerEntry){
	srvEntry := new(ServerEntry)
	srvEntry.ip_port = ip_port
	srvEntry.name = name
	srvEntry.publicKey = publicKey
	return srvEntry
}

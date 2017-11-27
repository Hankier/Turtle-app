package serverEntry

type ServerEntry struct{
	Name string
	Ip_port string
	PublicKey [256]byte
}

func NewServerEntry(name string, ip_port string, publicKey [256]byte)(*ServerEntry){
	srvEntry := new(ServerEntry)
	srvEntry.Ip_port = ip_port
	srvEntry.Name = name
	srvEntry.PublicKey = publicKey
	return srvEntry
}

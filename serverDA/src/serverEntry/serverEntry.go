package serverEntry

type ServerEntry struct{
	Name string
	Ip_port string
	PublicKey []byte
}

func NewServerEntry(name string, ip_port string, publicKey []byte)(*ServerEntry){
	srvEntry := new(ServerEntry)
	srvEntry.Ip_port = ip_port
	srvEntry.Name = name
	srvEntry.PublicKey = publicKey
	return srvEntry
}

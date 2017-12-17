package entry

type Entry struct{
	Name string
	Ip_port string
	PublicKey []byte
}

func NewEntry(name string, ip_port string, publicKey []byte)(*Entry){
	srvEntry := new(Entry)
	srvEntry.Ip_port = ip_port
	srvEntry.Name = name
	srvEntry.PublicKey = publicKey
	return srvEntry
}

package entry

type Entry struct{
	Name      string
	Ipport    string
	PublicKey []byte
}

func New(name string, ip_port string, publicKey []byte)(*Entry){
	srvEntry := new(Entry)
	srvEntry.Ipport = ip_port
	srvEntry.Name = name
	srvEntry.PublicKey = publicKey
	return srvEntry
}

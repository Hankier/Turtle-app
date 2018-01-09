package entry

//TODO change it!
type Entry struct{
	Name string
	Ip_port string
	PublicKey string
}

func NewEntry(name string, ip_port string, publicKey string)(*Entry){
	srvEntry := new(Entry)
	srvEntry.Ip_port = ip_port
	srvEntry.Name = name
	srvEntry.PublicKey = publicKey
	return srvEntry
}

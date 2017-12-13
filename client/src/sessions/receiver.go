package sessions

type Receiver interface{
	OnReceive(name string, content []byte)
}

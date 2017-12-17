package receiver

type Receiver interface{
	OnReceive(name string, content []byte)
}

package session

type Receiver interface{
	OnReceive(content []byte)
}

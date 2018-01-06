package receiver

type Receiver interface{
	OnReceive(from string, content []byte)
}
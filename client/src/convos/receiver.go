package convos

type Receiver interface{
	OnReceive(from string, content []byte)
}
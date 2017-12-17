package session

type Sender interface{
	Send(content []byte)
}

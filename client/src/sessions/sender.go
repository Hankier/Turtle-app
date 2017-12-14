package sessions

type Sender interface{
	Send(name string, content []byte)error
	SendInstant(name string, content []byte)error
	UnlockSending(name string)error
}

package session

type Sender interface{
	Send(content []byte)
	SendInstant(content []byte)
	UnlockSending()
}
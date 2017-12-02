package sender

type Sender interface{
	Send(bytes []byte)
	SendInstant(bytes []byte)
	UnlockSending()
}
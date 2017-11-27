package sender

type Sender interface{
	Send(bytes []byte)
	UnlockSending()
}
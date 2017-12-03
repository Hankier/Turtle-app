package sessionSender

type SessionSender interface{
	SendTo(bytes []byte)error
	SendInstantTo(bytes []byte)error
	UnlockSending()
}

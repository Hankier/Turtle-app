package sessionsSender

type SessionsSender interface{
	SendTo(name string, bytes []byte)
	UnlockSending(name string)
}

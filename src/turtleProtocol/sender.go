package turtleProtocol

type Sender interface{
	Send(content []byte)
	SendInstant(content []byte)
	UnlockSending()
}
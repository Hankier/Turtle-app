package turtleProtocol

type Receiver interface{
	OnReceive(content []byte)
}

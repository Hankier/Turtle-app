package receiver

import "turtleProtocol/msg"

type Receiver interface{
	OnReceive(name string, message *msg.Message)
}

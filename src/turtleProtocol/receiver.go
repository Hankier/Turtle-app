package turtleProtocol

import "turtleProtocol/msg"

type Receiver interface{
	OnReceive(message *msg.Message)
}

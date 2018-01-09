package sender

import "turtleProtocol/msg"

type Sender interface{
	Send(name string, message *msg.Message)error
}

package parser

import "turtleProtocol/msg"

type Parser interface{
	ParseMessage(from string, message *msg.Message)
}

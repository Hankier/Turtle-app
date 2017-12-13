package session

import (
	"message"
	"utils"
)

type Sender interface{
	Send(content []byte)
	SendInstant(content []byte)
	UnlockSending()
}
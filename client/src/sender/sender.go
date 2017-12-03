package sender

import "message"

type Sender interface{
	Send(msg message.Message)
	SendInstant(msg message.Message)
	UnlockSending()
}
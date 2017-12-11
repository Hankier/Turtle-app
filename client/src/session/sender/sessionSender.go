package sender

import "message"

type Sender interface{
	Send(msg *message.Message)error
	SendInstant(msg *message.Message)error
	UnlockSending()
}

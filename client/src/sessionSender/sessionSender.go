package sessionSender

import "message"

type SessionSender interface{
	Send(msg *message.Message)error
	SendInstant(msg *message.Message)error
	UnlockSending()
}

package sessionSender

import "message"

type SessionSender interface{
	SendTo(msg message.Message)error
	SendInstantTo(msg message.Message)error
	UnlockSending()
}

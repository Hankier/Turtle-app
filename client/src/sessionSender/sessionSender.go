package sessionSender

import "message"

type SessionSender interface{
	Send(content []byte, receiver string, receiverServer string)error
	SendInstant(msg *message.Message)error
	UnlockSending()
}

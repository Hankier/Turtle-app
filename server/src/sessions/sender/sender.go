package sender

import "message"

type Sender interface{
	SendTo(name string, msg *message.Message)
	SendInstantTo(name string, msg *message.Message)
	UnlockSending(name string)
}

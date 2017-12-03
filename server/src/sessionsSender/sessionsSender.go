package sessionsSender

import "message"

type SessionsSender interface{
	SendTo(name string, msg *message.Message)
	SendInstantTo(name string, msg *message.Message)
	UnlockSending(name string)
}

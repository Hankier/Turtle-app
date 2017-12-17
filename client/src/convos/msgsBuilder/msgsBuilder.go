package msgsBuilder

import "crypt"

type MessageBuilder interface{
	BuildMessageContent(server string, name string, command string, encType crypt.TYPE)([]byte, error)
}

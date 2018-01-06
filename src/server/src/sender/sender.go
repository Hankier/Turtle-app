package sender

import (
	"message"
	"utils"
)

type Sender interface{
	Send(msg *message.Message)
	SendInstant(msg *message.Message)
	UnlockSending()
}

func addSizeToBytes(bytes []byte)([]byte){
	size := utils.IntToTwobytes(len(bytes))

	bytes = append(size, bytes...)

	return bytes
}
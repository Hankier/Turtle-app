package messageHandler

type MessageHandler interface{
	HandleBytes(bytes []byte)
}


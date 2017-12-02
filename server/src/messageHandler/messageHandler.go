package messageHandler

type MessageHandler interface{
	HandleBytes(from string, bytes []byte)
}


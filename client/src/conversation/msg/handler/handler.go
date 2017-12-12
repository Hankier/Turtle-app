package handler

type ConversationMessageHandler interface{
	HandleBytes(from string, bytes []byte)
}

package conversationMessageHandler

type ConversationMessageHandler interface{
	HandleBytes(from string, bytes []byte)
}

package conversationMessageHandler

type ConversationMessageHandler interface{
	HandleBytes(bytes []byte)
}

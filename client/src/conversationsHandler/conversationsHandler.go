package conversationsHandler

type ConversationsHandler interface{
	ReceiveMessage(content []byte, receiver string, receiverServer string)error
}

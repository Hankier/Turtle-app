package userInterface

import "conversation"

type UserInterface interface{
	GetCurrentPath()([]string)
	ChooseNewPath(length int)([]string)
	ConnectToServer(name string)error
	GetServerList()([]string)
	CreateConversation(receiver string, receiverServer string) (convo *conversation.Conversation, err error)
	SendTo(message string, receiver string, receiverServer string)error
}

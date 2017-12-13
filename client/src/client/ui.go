package client

import "convos"

type UserInterface interface{
	GetCurrentPath()([]string)
	ChooseNewPath(length int)([]string,error)
	ConnectToServer(name string)error
	GetServerList()([]string)
	CreateConversation(receiver string, receiverServer string) (convo *convos.Conversation, err error)
	SendTo(message string, receiver string, receiverServer string)error
}

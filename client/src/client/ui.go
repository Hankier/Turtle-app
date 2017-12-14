package client

type UserInterface interface{
	GetCurrentPath()([]string)
	ChooseNewPath(length int)([]string,error)
	ConnectToServer(name string)error
	GetServerList()([]string)
	CreateConversation(receiver string, receiverServer string) (err error)
	SendTo(message string, receiverServer string, receiver string)error
}

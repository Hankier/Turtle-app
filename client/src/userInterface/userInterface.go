package userInterface


type UserInterface interface{
	GetCurrentPath()([]string)
	ChooseNewPath()([]string)
	ConnectToServer(name string)
	GetServerList()([]string)
	SendTo(message, receiver, recServer string)
}

package userInterface

import "serverEntry"

type UserInterface interface{
	GetCurrentPath()([]serverEntry.ServerEntry)
	ChooseNewPath()([]serverEntry.ServerEntry)
	ConnectToServer(name string)
	GetServerList()([]serverEntry.ServerEntry)
	SendTo(message, receiver, recServer string)
}

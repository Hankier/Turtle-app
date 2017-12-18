package ui

import "crypt"

//Interface enabling interaction with user
//
//GetCurrentPath returns a slice of strings representing consecutive nodes of path
//
//ChooseNewPath generates a new, random path of length given as a parameter and assigning it as a current path
//Returns a slice of strings representing consecutive nodes of path and an error accordingly to encountered problems, nil if all went well
//
//ConnectToSever tries to connect to a server with a given server name
//Returns an error accordingly to encountered problems, nil if all went well
//
//GetServerList returns a slice containing all known server names as strings
//
//CreateConversation tries to create a conversation with given receiver which should be connected to given receiverServer
//Returns an error accordingly to encountered problems, nil if all went well
//
//SendTo tries to send a given message to a receiver connected to a given receiverServer
//Returns an error accordingly to encountered problems, nil if all went well
type UserInterface interface{

	GetCurrentPath()([]string)
	ChooseNewPath(length int)([]string,error)
	SetEncryptionType(enctype crypt.TYPE)

	ConnectToServer(name string)error

	GetServerList()([]string)
	GetServerDetails(name string)([]string)

	CreateConversation(receiverServer string, receiver string) (err error)
	SetConversationKey(receiverServer string, receiver string, enctype crypt.TYPE, filename string) error
	SendTo(receiverServer string, receiver string, message string)error
}

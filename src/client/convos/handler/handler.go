package handler

import "crypt"

type Handler interface{
	CreateConversation(server string, name string)
	RemoveConversation(server string, name string)
	GetActiveConversations()(server string, name string)
	SetConversationKey(server string, name string, enctype crypt.TYPE, keydata []byte)error
}

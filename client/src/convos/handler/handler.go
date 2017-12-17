package handler

type Handler interface{
	CreateConversation(server string, name string)
	RemoveConversation(server string, name string)
	GetActiveConversations()(server string, name string)
}

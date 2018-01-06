package credentials

//Interface holding its own name and a server it is currently connected to
type CredentialsHolder interface{
	GetName()string
	GetCurrentServer()(string, error)
}

package client

type CredentialsHolder interface{
	GetName()string
	GetCurrentServer()string
}

package dialer

type Dialer interface{
	ConnectToServer(name string)error
}
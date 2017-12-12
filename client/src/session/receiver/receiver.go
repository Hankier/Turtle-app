package receiver

type Receiver interface{
	onReceive([]byte)
}

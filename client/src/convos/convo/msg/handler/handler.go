package handler

type Handler interface{
	HandleBytes(from string, bytes []byte)
}

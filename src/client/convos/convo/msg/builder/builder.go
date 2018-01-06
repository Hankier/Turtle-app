package builder

type Builder interface{
	Build()[]byte
	ParseCommand(message string)
}
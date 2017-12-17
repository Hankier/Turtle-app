package parser

type Parser interface{
	ParseBytes(from string, bytes []byte)
}


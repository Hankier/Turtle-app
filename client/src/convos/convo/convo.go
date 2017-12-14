package convo

import(
	"textReceiver"
	"convos/convo/msg/handler"
	"convos/convo/msg/builder"
	"convos/convo/key"
	"convos/convo/encrypter"
)

type Conversation struct{
	name                string
	server              string
	commonKey   		key.CommonKey
	encrypter   		encrypter.Encrypter
	textReceiver        textReceiver.TextReceiver
	convoMessageBuilder *builder.BuilderImpl
	convoMessageHandler handler.Handler
}

func New(textReceiver textReceiver.TextReceiver, name string, server string)*Conversation{
	convo := new(Conversation)
	convo.name = name
	convo.server = server
	convo.commonKey = key.New()
	convo.encrypter = encrypter.New()
	convo.textReceiver = textReceiver
	convo.convoMessageBuilder = builder.New(convo.commonKey)
	convo.convoMessageHandler = handler.New(convo.commonKey, convo.encrypter, convo.textReceiver)
	return convo
}

func (convo *Conversation)Receive(msg []byte){
	convo.convoMessageHandler.HandleBytes(convo.name + " " + convo.server, msg)
}

func (convo *Conversation)BuildMessageContent(command string)[]byte{
	convo.convoMessageBuilder.ParseCommand(command)
	return convo.convoMessageBuilder.Build()
}
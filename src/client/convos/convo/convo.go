package convo

import(
	"client/textReceiver"
	"client/convos/convo/msg/handler"
	"client/convos/convo/msg/builder"
	"client/convos/convo/key"
	"client/convos/convo/encrypter"
	"client/client/credentials"
	"crypt"
)

type Conversation struct{
	name                string
	server              string
	commonKey    		key.CommonKey
	encrypter    		encrypter.Encrypter
	textReceiver 		textReceiver.TextReceiver
	msgBuilder   		*builder.BuilderImpl
	msgHandler   		handler.Handler
	credHolder   		credentials.CredentialsHolder
}

func New(name string, server string, textReceiver textReceiver.TextReceiver, holder credentials.CredentialsHolder)*Conversation{
	convo := new(Conversation)
	convo.name = name
	convo.server = server
	convo.commonKey = key.New()
	convo.encrypter = encrypter.New()
	convo.textReceiver = textReceiver
	convo.msgBuilder = builder.New(convo.commonKey)
	convo.msgHandler = handler.New(convo.commonKey, convo.encrypter, convo.textReceiver)
	convo.credHolder = holder
	return convo
}

func (convo *Conversation)Receive(msg []byte){
	convo.msgHandler.HandleBytes(convo.name + " " + convo.server, msg)
}

func (convo *Conversation)SetKey(enctype crypt.TYPE, keydata []byte)error{
	return convo.encrypter.SetKey(enctype, keydata)
}

func (convo *Conversation)BuildMessageContent(command string, encType crypt.TYPE)[]byte{
	convo.msgBuilder.ParseCommand(command)
	//TODO error handlign
	server, _ := convo.credHolder.GetCurrentServer()
	name := convo.credHolder.GetName()
	content := ([]byte)(server + name)
	content = append(content, convo.msgBuilder.Build()...)
	//TODO error handling
	content, _ = convo.encrypter.Encrypt(encType, content)
	return content
}
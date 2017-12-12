package conversation

import(
	"commonKeyProtocol"
	"receiverEncrypter"
	"textReceiver"
	"conversation/msg/handler"
	"conversation/msg/builder"
	"crypt"
)

type Conversation struct{
	name                string
	server              string
	commonKeyProtocol   commonKeyProtocol.CommonKeyProtocol
	receiverEncrypter   *receiverEncrypter.ReceiverEncrypterImpl
	textReceiver        textReceiver.TextReceiver
	convoMessageBuilder *builder.ConversationMessageBuilder
	convoMessageHandler handler.ConversationMessageHandler
}

func New(textReceiver textReceiver.TextReceiver, name string, server string)*Conversation{
	convo := new(Conversation)
	convo.name = name
	convo.server = server
	convo.commonKeyProtocol = commonKeyProtocol.New()
	convo.receiverEncrypter = receiverEncrypter.New()
	convo.textReceiver = textReceiver
	convo.convoMessageBuilder = builder.New(convo.commonKeyProtocol)
	convo.convoMessageHandler = handler.New(convo.commonKeyProtocol, convo.receiverEncrypter, convo.textReceiver)
	return convo
}

func (convo *Conversation)Receive(msg []byte){
	convo.convoMessageHandler.HandleBytes(convo.name + " " + convo.server, msg)
}

func (convo *Conversation)MessageBuilder()*builder.ConversationMessageBuilder{
	return convo.convoMessageBuilder
}

func (convo *Conversation) Encrypter() crypt.Encrypter{
	return convo.receiverEncrypter
}
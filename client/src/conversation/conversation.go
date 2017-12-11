package conversation

import(
	"commonKeyProtocol"
	"receiverEncrypter"
	"textReceiver"
	"conversationMessageHandler"
	"conversationMessageBuilder"
	"cryptographer"
)

type Conversation struct{
	name                string
	server              string
	commonKeyProtocol   commonKeyProtocol.CommonKeyProtocol
	receiverEncrypter   *receiverEncrypter.ReceiverEncrypterImpl
	textReceiver        textReceiver.TextReceiver
	convoMessageBuilder *conversationMessageBuilder.ConversationMessageBuilder
	convoMessageHandler conversationMessageHandler.ConversationMessageHandler
}

func NewConversation(textReceiver textReceiver.TextReceiver, name string, server string)*Conversation{
	convo := new(Conversation)
	convo.name = name
	convo.server = server
	convo.commonKeyProtocol = commonKeyProtocol.New()
	convo.receiverEncrypter = receiverEncrypter.New()
	convo.textReceiver = textReceiver
	convo.convoMessageBuilder = conversationMessageBuilder.NewConversationMessageBuilder(convo.commonKeyProtocol)
	convo.convoMessageHandler = conversationMessageHandler.New(convo.commonKeyProtocol, convo.receiverEncrypter, convo.textReceiver)
	return convo
}

func (convo *Conversation)Receive(msg []byte){
	convo.convoMessageHandler.HandleBytes(convo.name + " " + convo.server, msg)
}

func (convo *Conversation)MessageBuilder()*conversationMessageBuilder.ConversationMessageBuilder{
	return convo.convoMessageBuilder
}

func (convo *Conversation) Encrypter()cryptographer.Encrypter{
	return convo.receiverEncrypter
}
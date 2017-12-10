package conversation

import(
	"commonKeyProtocol"
	"receiverKeyHandler"
	"textReceiver"
	"conversationMessageHandler"
)

type Conversation struct{
	name string
	server string
	conversationMessageBuilder *conversationMessageBuilder.ConversationMessageBuilder
	conversationMessageHandler conversationMessageHandler.ConversationMessageHandler
	commonKeyProtocol commonKeyProtocol.CommonKeyProtocol
	receiverKeyHandler receiverKeyHandler.ReceiverKeyHandler
	textReceiver textReceiver.TextReceiver
}

func NewConversation(textReceiver textReceiver.TextReceiver, name string, server string){
	convo := new(Conversation)
	convo.name = name
	convo.server = server
	convo.textReceiver = textReceiver
	convo.commonKeyProtocol = commonKeyProtocol.NewCommonKeyProtocol()
	convo.receiverKeyHandler = receiverKeyHandler.NewReceiverKeyHandler()
}

func (convo *Conversation)Receive(msg []byte){
	convo.conversationMessageHandler.HandleBytes(msg)
}

func (convo *Conversation)MessageBuilder()*conversationMessageBuilder.ConversationMessageBuilder{
	return convo.conversationMessageBuilder
}
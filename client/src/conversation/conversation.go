package conversation

import(
	"commonKeyProtocol"
	"receiverKeyHandler"
	"textReceiver"
	"conversationMessageHandler"
	"conversationMessageBuilder"
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

func NewConversation(textReceiver textReceiver.TextReceiver, name string, server string)*Conversation{
	convo := new(Conversation)
	convo.name = name
	convo.server = server
	convo.textReceiver = textReceiver
	convo.commonKeyProtocol = commonKeyProtocol.NewCommonKeyProtocolImpl()
	convo.receiverKeyHandler = receiverKeyHandler.NewReceiverKeyHandlerImpl()
	return convo
}

func (convo *Conversation)Receive(msg []byte){
	convo.conversationMessageHandler.HandleBytes(msg)
}

func (convo *Conversation)MessageBuilder()*conversationMessageBuilder.ConversationMessageBuilder{
	return convo.conversationMessageBuilder
}

func (convo *Conversation)ReceiverKeyHandler()receiverKeyHandler.ReceiverKeyHandler{
	return convo.receiverKeyHandler
}
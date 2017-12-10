package conversation

import(
	"commonKeyProtocol"
	"receiverKeyHandler"
	"textReceiver"
	"conversationMessageHandler"
	"conversationMessageBuilder"
)

type Conversation struct{
	name                string
	server              string
	commonKeyProtocol   commonKeyProtocol.CommonKeyProtocol
	receiverKeyHandler  receiverKeyHandler.ReceiverKeyHandler
	textReceiver        textReceiver.TextReceiver
	convoMessageBuilder *conversationMessageBuilder.ConversationMessageBuilder
	convoMessageHandler conversationMessageHandler.ConversationMessageHandler
}

func NewConversation(textReceiver textReceiver.TextReceiver, name string, server string)*Conversation{
	convo := new(Conversation)
	convo.name = name
	convo.server = server
	convo.commonKeyProtocol = commonKeyProtocol.NewCommonKeyProtocolImpl()
	convo.receiverKeyHandler = receiverKeyHandler.NewReceiverKeyHandlerImpl()
	convo.textReceiver = textReceiver
	convo.convoMessageBuilder = conversationMessageBuilder.NewConversationMessageBuilder(convo.commonKeyProtocol)
	convo.convoMessageHandler = conversationMessageHandler.NewConversationMessageHandlerImpl(convo.commonKeyProtocol, convo.receiverKeyHandler, convo.textReceiver)
	return convo
}

func (convo *Conversation)Receive(msg []byte){
	convo.convoMessageHandler.HandleBytes(convo.name + " " + convo.server, msg)
}

func (convo *Conversation)MessageBuilder()*conversationMessageBuilder.ConversationMessageBuilder{
	return convo.convoMessageBuilder
}

func (convo *Conversation)ReceiverKeyHandler()receiverKeyHandler.ReceiverKeyHandler{
	return convo.receiverKeyHandler
}
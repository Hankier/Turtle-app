package conversationMessageHandler

func (convMHI *ConversationMessageHandlerImpl)handleDEFAULT(from string, bytes []byte){
	convMHI.textRecv.Print(from, string(bytes))
}

func (convMHI *ConversationMessageHandlerImpl)handleCOMMON_KEY_PROTOCOL(bytes []byte){

}

func (convMHI *ConversationMessageHandlerImpl)handleINIT_DATA(bytes []byte){

}
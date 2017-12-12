package handler

func (convMHI *ConversationMessageHandlerImpl)handleDEFAULT(from string, bytes []byte){
	convMHI.textrecv.Print(from, string(bytes))
}

func (convMHI *ConversationMessageHandlerImpl)handleCOMMON_KEY_PROTOCOL(bytes []byte){

}

func (convMHI *ConversationMessageHandlerImpl)handleINIT_DATA(bytes []byte){

}
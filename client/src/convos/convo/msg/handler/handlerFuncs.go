package handler

func (convMHI *HandlerImpl)handleDEFAULT(from string, bytes []byte){
	convMHI.textrecv.Print(from, string(bytes))
}

func (convMHI *HandlerImpl)handleCOMMON_KEY_PROTOCOL(bytes []byte){

}

func (convMHI *HandlerImpl)handleINIT_DATA(bytes []byte){

}
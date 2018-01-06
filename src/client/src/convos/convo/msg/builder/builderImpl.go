package builder

import (
	"convos/convo/msg"
	"convos/convo/key"
)

type BuilderImpl struct{
	commonKey      key.CommonKey
	messageType    msg.TYPE
	encType        key.TYPE
	messageContent []byte
}

func New(commonKey key.CommonKey)*BuilderImpl {
	builder := new(BuilderImpl)
	builder.commonKey = commonKey
	builder.messageType = msg.DEFAULT
	builder.encType = key.PLAIN
	builder.messageContent = make([]byte, 0, 0)
	return builder
}

func (builder *BuilderImpl)SetEncryption(encType key.TYPE){
	builder.encType = encType
}

func (builder *BuilderImpl)SetMessage(message string){
	builder.messageContent = ([]byte)(message)
}

func (builder *BuilderImpl)SetCommonKeyData(part int, content []byte){
	builder.messageContent = builder.commonKey.GetCommonKeyData(part)
}

func (builder *BuilderImpl)SetInitData(){
	//TODO get this data from nodeCrypto
}

func (builder *BuilderImpl)Build()[]byte{
	convoMsg := msg.New(builder.messageType, builder.encType, builder.messageContent)
	return convoMsg.ToBytes()
}

func (builder *BuilderImpl)ParseCommand(message string){
	builder.SetMessage(message)
	//TODO
}

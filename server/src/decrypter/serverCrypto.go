package decrypter

type ServerCrypto struct{

}

func NewServerCrypto()(*ServerCrypto){
	return new(ServerCrypto);
}

func (crypto *ServerCrypto)Decrypt(bytes []byte)[]byte{
	return bytes
}
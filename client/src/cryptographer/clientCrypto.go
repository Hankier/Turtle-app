package cryptographer

import "crypto/rsa"

type ClientCrypto struct{
	privateKey  *rsa.PrivateKey
	publicKey   *rsa.PublicKey
}


func NewClientCrypto()(*ClientCrypto){
	cli := new(ClientCrypto)
	return cli
}

func (cli *ClientCrypto)encrypt(bytes []byte)[]byte{
	return bytes
}

func (cli *ClientCrypto)decrypt(bytes []byte)[]byte{
	return bytes
}
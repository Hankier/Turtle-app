package cryptographer

import (
	"crypto/rsa"
)

type ClientCrypto struct{
	privateKey  *rsa.PrivateKey
	publicKey   *rsa.PublicKey
}

func NewClientCrypto()(*ClientCrypto){
	cli := new(ClientCrypto)
	return cli
}

func (cli *ClientCrypto)Encrypt(encType TYPE, bytes []byte)[]byte{
	switch encType {
	case PLAIN:
		return EncryptPlain(bytes)
	case RSA:
		return EncryptRSA(cli.publicKey, bytes)
	case ELGAMAL:
		return EncryptElGamal(bytes)
	}
	return bytes
}

func (cli *ClientCrypto)Decrypt(encType TYPE, bytes []byte)[]byte{
	switch encType {
	case PLAIN:
		return DecryptPlain(bytes)
	case RSA:
		return DecryptRSA(cli.privateKey, bytes)
	case ELGAMAL:
		return DecryptElGamal(bytes)
	}
	return bytes
}
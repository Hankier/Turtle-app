package cryptographer

import (
	"crypto/rsa"
	"golang.org/x/crypto/openpgp/elgamal"
)

type NodeCrypto struct{
	privateKeyRSA  *rsa.PrivateKey
	publicKeyRSA   *rsa.PublicKey
	privateKeyElGamal  *elgamal.PrivateKey
	publicKeyElGamal   *elgamal.PublicKey
}

func NewNodeCrypto()(*NodeCrypto){
	nodeCrypto := new(NodeCrypto)

	var privateKeyRSA *rsa.PrivateKey

	privateKeyRSA, err := LoadRSA("privateKeyRSA")
	if err != nil{
		privateKeyRSA = GenerateRSA()
		SaveRSA(privateKeyRSA, "privateKeyRSA")
	}
	nodeCrypto.privateKeyRSA = privateKeyRSA
	nodeCrypto.publicKeyRSA = &privateKeyRSA.PublicKey


	var privateKeyElGamal *elgamal.PrivateKey

	privateKeyElGamal, err = LoadElGamal("privateKeyElGamal")
	if err != nil{
		privateKeyElGamal = GenerateElGamal()
		SaveElGamal(privateKeyElGamal, "privateKeyElGamal")
	}
	nodeCrypto.privateKeyElGamal = privateKeyElGamal
	nodeCrypto.publicKeyElGamal = &privateKeyElGamal.PublicKey

	return nodeCrypto
}

func (nodeCrypto *NodeCrypto)Decrypt(encType TYPE, bytes []byte) ([]byte, error){
	switch encType {
	case PLAIN:
		return DecryptPlain(bytes), nil
	case RSA:
		return DecryptRSA(nodeCrypto.privateKeyRSA, bytes)
	case ELGAMAL:
		return DecryptElGamal(nodeCrypto.privateKeyElGamal, bytes)
	}
	return bytes, nil
}

func (nodeCrypto *NodeCrypto)Encrypt(encType TYPE, bytes []byte) ([]byte, error){
	switch encType {
	case PLAIN:
		return EncryptPlain(bytes), nil
	case RSA:
		return EncryptRSA(nodeCrypto.publicKeyRSA, bytes)
	case ELGAMAL:
		return EncryptElGamal(nodeCrypto.publicKeyElGamal, bytes)
	}
	return bytes, nil
}

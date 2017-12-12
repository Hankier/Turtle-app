package crypt

import (
	"crypto/rsa"
	"golang.org/x/crypto/openpgp/elgamal"
)

type NodeCrypto struct{
	privRSA           *rsa.PrivateKey
	pubRSA            *rsa.PublicKey
	privElGamal *elgamal.PrivateKey
	pubElGamal        *elgamal.PublicKey
}

func New()(*NodeCrypto){
	nc := new(NodeCrypto)

	var privRSA *rsa.PrivateKey

	privRSA, err := LoadRSA("privateKeyRSA")
	if err != nil{
		privRSA = GenerateRSA()
		SaveRSA(privRSA, "privateKeyRSA")
	}
	nc.privRSA = privRSA
	nc.pubRSA = &privRSA.PublicKey


	var privElGamal *elgamal.PrivateKey

	privElGamal, err = LoadElGamal("privateKeyElGamal")
	if err != nil{
		privElGamal = GenerateElGamal()
		SaveElGamal(privElGamal, "privateKeyElGamal")
	}
	nc.privElGamal = privElGamal
	nc.pubElGamal = &privElGamal.PublicKey

	return nc
}

func (nc *NodeCrypto)Decrypt(encType TYPE, bytes []byte) ([]byte, error){
	switch encType {
	case PLAIN:
		return bytes, nil
	case RSA:
		return DecryptRSA(nc.privRSA, bytes)
	case ELGAMAL:
		return DecryptElGamal(nc.privElGamal, bytes)
	}
	return bytes, nil
}

func (nc *NodeCrypto)Encrypt(encType TYPE, bytes []byte) ([]byte, error){
	switch encType {
	case PLAIN:
		return bytes, nil
	case RSA:
		return EncryptRSA(nc.pubRSA, bytes)
	case ELGAMAL:
		return EncryptElGamal(nc.pubElGamal, bytes)
	}
	return bytes, nil
}

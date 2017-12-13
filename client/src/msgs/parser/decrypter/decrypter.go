package decrypter

import (
	"crypt"
	"crypto/rsa"
	"golang.org/x/crypto/openpgp/elgamal"
)

type ClientDecrypterImpl struct{
	privRSA           *rsa.PrivateKey
	privElGamal *elgamal.PrivateKey
}

func New()(*ClientDecrypterImpl){
	nc := new(ClientDecrypterImpl)

	var privRSA *rsa.PrivateKey

	privRSA, err := crypt.LoadRSA("privateKeyRSA")
	if err != nil{
		privRSA = crypt.GenerateRSA()
		crypt.SaveRSA(privRSA, "privateKeyRSA")
	}
	nc.privRSA = privRSA


	var privElGamal *elgamal.PrivateKey

	privElGamal, err = crypt.LoadElGamal("privateKeyElGamal")
	if err != nil{
		privElGamal = crypt.GenerateElGamal()
		crypt.SaveElGamal(privElGamal, "privateKeyElGamal")
	}
	nc.privElGamal = privElGamal

	return nc
}

func (nc *ClientDecrypterImpl)Decrypt(encType crypt.TYPE, bytes []byte) ([]byte, error){
	switch encType {
	case crypt.PLAIN:
		return bytes, nil
	case crypt.RSA:
		return crypt.DecryptRSA(nc.privRSA, bytes)
	case crypt.ELGAMAL:
		return crypt.DecryptElGamal(nc.privElGamal, bytes)
	}
	return bytes, nil
}

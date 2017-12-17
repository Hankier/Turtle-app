package decrypter

import (
	"crypt"
	"crypto/rsa"
	"golang.org/x/crypto/openpgp/elgamal"
)

type DecrypterImpl struct{
	privRSA           *rsa.PrivateKey
	privElGamal *elgamal.PrivateKey
}

func New()(*DecrypterImpl){
	nc := new(DecrypterImpl)

	var privRSA *rsa.PrivateKey

	privRSA, err := crypt.LoadRSA("privateKeyRSA")
	if err != nil{
		privRSA = crypt.GenerateRSA()
		crypt.SaveRSA(privRSA, "privateKeyRSA")
		crypt.SaveRSAPublic(&privRSA.PublicKey, "publicKeyRSA")
	}
	nc.privRSA = privRSA


	var privElGamal *elgamal.PrivateKey

	privElGamal, err = crypt.LoadElGamal("privateKeyElGamal")
	if err != nil{
		privElGamal = crypt.GenerateElGamal()
		crypt.SaveElGamal(privElGamal, "privateKeyElGamal")
		//TODO saving public key
	}
	nc.privElGamal = privElGamal

	return nc
}

func (nc *DecrypterImpl)Decrypt(encType crypt.TYPE, bytes []byte) ([]byte, error){
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

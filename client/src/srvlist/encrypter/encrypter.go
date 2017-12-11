package encrypter

import (
	"cryptographer"
	"crypto/rsa"
	"golang.org/x/crypto/openpgp/elgamal"
)

type ServerEncrypterImpl struct{
	pubRSA     *rsa.PublicKey
	pubElGamal *elgamal.PublicKey
}

func New(pubRSA *rsa.PublicKey, pubElGamal *elgamal.PublicKey)(*ServerEncrypterImpl){
	enc := new(ServerEncrypterImpl)
	enc.pubElGamal = pubElGamal
	enc.pubRSA = pubRSA
	return enc
}

func (enc *ServerEncrypterImpl)Encrypt(enctype cryptographer.TYPE, msg []byte)([]byte, error){
	switch enctype {
	case cryptographer.PLAIN:
		return msg, nil
	case cryptographer.RSA:
		return cryptographer.EncryptRSA(enc.pubRSA, msg)
	case cryptographer.ELGAMAL:
		return cryptographer.EncryptElGamal(enc.pubElGamal, msg)
	}

	return msg, nil
}

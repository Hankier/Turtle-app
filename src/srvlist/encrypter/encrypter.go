package encrypter

import (
	"crypt"
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

func (enc *ServerEncrypterImpl)GetPublicKeyRSA() *rsa.PublicKey{
	return enc.pubRSA;
}

func (enc *ServerEncrypterImpl)GetPublicKeyElGamal() *elgamal.PublicKey{
	return enc.pubElGamal;
}

func (enc *ServerEncrypterImpl)Encrypt(enctype crypt.TYPE, msg []byte)([]byte, error){
	switch enctype {
	case crypt.PLAIN:
		return msg, nil
	case crypt.RSA:
		return crypt.EncryptRSA(enc.pubRSA, msg)
	case crypt.ELGAMAL:
		return crypt.EncryptElGamal(enc.pubElGamal, msg)
	}

	return msg, nil
}

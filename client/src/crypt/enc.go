package crypt

import (
	"crypto/rsa"
	"golang.org/x/crypto/openpgp/elgamal"
)

type Encrypter interface{
	Encrypt(encType TYPE, bytes []byte) ([]byte, error)
	GetPublicKeyRSA() *rsa.PublicKey
	GetPublicKeyElGamal() *elgamal.PublicKey
}

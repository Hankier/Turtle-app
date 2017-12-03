package cryptographer

import (
	"crypto/rsa"
	"log"
	"crypto/rand"
)

type TYPE byte

const (
	PLAIN TYPE = iota
	RSA
	ELGAMAL
)

type Cryptographer interface{
	Encrypt(encType TYPE, bytes []byte)[]byte
	Decrypt(encType TYPE, bytes []byte)[]byte
}

func DecryptRSA(privateKey *rsa.PrivateKey, msg []byte)[]byte{
	decryptedText, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, msg);
	if  err != nil {
		log.Fatal(err)
	}
	return decryptedText
}

func DecryptElGamal(bytes []byte) []byte {
	return bytes
	//TODO ELGAMAL!!!
}

func DecryptPlain(bytes []byte) []byte {
	return bytes
}

func EncryptRSA(publicKey *rsa.PublicKey, msg []byte)[]byte{
	encryptedText, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, msg);
	if  err != nil {
		log.Fatal(err)
	}
	return encryptedText
}

func EncryptElGamal(bytes []byte) []byte {
	return bytes
	//TODO ELGAMAL!!!
}

func EncryptPlain(bytes []byte) []byte {
	return bytes
}

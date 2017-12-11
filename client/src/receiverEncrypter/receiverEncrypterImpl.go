package receiverEncrypter

import "crypto/rsa"
import (
	"golang.org/x/crypto/openpgp/elgamal"
	"cryptographer"
	"encoding/pem"
	"math/big"
	"crypto/x509"
	"log"
)

type ReceiverEncrypterImpl struct{
	pubRSA     *rsa.PublicKey
	pubElGamal *elgamal.PublicKey
}

func New()(*ReceiverEncrypterImpl){
	return new(ReceiverEncrypterImpl)
}

func (recv *ReceiverEncrypterImpl)SetKey(enctype cryptographer.TYPE, keydata []byte){
	switch enctype {
	case cryptographer.RSA:
		recv.setRSA(keydata)
		break
	case cryptographer.ELGAMAL:
		recv.setElGamal(keydata)
		break
	}
}

func (recv *ReceiverEncrypterImpl)setRSA(keydata []byte){
	//todo public key
	block, _ := pem.Decode(keydata)

	pubRSA, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		log.Print("Failed to parse private key: " + err.Error())
		return
	}
	pubRSA.Precompute()

	if err = pubRSA.Validate(); err != nil {
		log.Print(err)
		return
	}
}

func (recv *ReceiverEncrypterImpl)setElGamal(keyData []byte){
	publicKeyElGamal := &elgamal.PublicKey{}

	block, keyData := pem.Decode(keyData)
	publicKeyElGamal.G = new(big.Int).SetBytes(block.Bytes)

	block, keyData = pem.Decode(keyData)
	publicKeyElGamal.P = new(big.Int).SetBytes(block.Bytes)

	block, _ = pem.Decode(keyData)
	publicKeyElGamal.Y = new(big.Int).SetBytes(block.Bytes)

	recv.pubElGamal = publicKeyElGamal
}

func (recv *ReceiverEncrypterImpl)Encrypt(encType cryptographer.TYPE, msg []byte)([]byte, error){
	switch encType {
	case cryptographer.PLAIN:
		return msg, nil
	case cryptographer.RSA:
		return cryptographer.EncryptRSA(recv.pubRSA, msg)
	case cryptographer.ELGAMAL:
		return cryptographer.EncryptElGamal(recv.pubElGamal, msg)
	}

	return msg, nil
}
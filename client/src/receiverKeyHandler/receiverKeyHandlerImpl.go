package receiverKeyHandler

import "crypto/rsa"
import (
	"golang.org/x/crypto/openpgp/elgamal"
	"cryptographer"
	"encoding/pem"
	"math/big"
	"crypto/x509"
	"log"
)

type ReceiverKeyHandlerImpl struct{
	publicKeyRSA   *rsa.PublicKey
	publicKeyElGamal *elgamal.PublicKey
}

func (recv *ReceiverKeyHandlerImpl)SetKey(encType cryptographer.TYPE, keyData []byte){
	switch encType{
	case cryptographer.RSA:
		recv.setRSA(keyData)
		break
	case cryptographer.ELGAMAL:
		recv.setElGamal(keyData)
		break
	}
}

func (recv *ReceiverKeyHandlerImpl)setRSA(keyData []byte){
	block, _ := pem.Decode(keyData)

	privateKeyRSA, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		log.Print("Failed to parse private key: " + err.Error())
		return
	}
	privateKeyRSA.Precompute()

	if err = privateKeyRSA.Validate(); err != nil {
		log.Print(err)
		return
	}
}

func (recv *ReceiverKeyHandlerImpl)setElGamal(keyData []byte){
	publicKeyElGamal := &elgamal.PublicKey{}

	block, keyData := pem.Decode(keyData)
	publicKeyElGamal.G = new(big.Int).SetBytes(block.Bytes)

	block, keyData = pem.Decode(keyData)
	publicKeyElGamal.P = new(big.Int).SetBytes(block.Bytes)

	block, _ = pem.Decode(keyData)
	publicKeyElGamal.Y = new(big.Int).SetBytes(block.Bytes)

	recv.publicKeyElGamal = publicKeyElGamal
}

func (recv *ReceiverKeyHandlerImpl)Encrypt(encType cryptographer.TYPE, msg []byte)([]byte, error){
	switch encType {
	case cryptographer.PLAIN:
		return cryptographer.EncryptPlain(msg), nil
	case cryptographer.RSA:
		return cryptographer.EncryptRSA(recv.publicKeyRSA, msg)
	case cryptographer.ELGAMAL:
		return cryptographer.EncryptElGamal(recv.publicKeyElGamal, msg)
	}

	return msg, nil
}
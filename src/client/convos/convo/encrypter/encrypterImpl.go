package encrypter

import "crypto/rsa"
import (
	"golang.org/x/crypto/openpgp/elgamal"
	"crypt"
	"encoding/pem"
	"math/big"
	"crypto/x509"
	"errors"
	"reflect"
)

type EncrypterImpl struct{
	pubRSA     *rsa.PublicKey
	pubElGamal *elgamal.PublicKey
}

func New()(*EncrypterImpl){
	return new(EncrypterImpl)
}

func (recv *EncrypterImpl)SetKey(enctype crypt.TYPE, keydata []byte)error{
	switch enctype {
	case crypt.RSA:
		return recv.setRSA(keydata)
		break
	case crypt.ELGAMAL:
		return recv.setElGamal(keydata)
		break
	}
	return nil
}

func (recv *EncrypterImpl)setRSA(keydata []byte)error{
	block, _ := pem.Decode(keydata)

	pubRSA, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return errors.New(reflect.TypeOf(recv).String() + ": " + err.Error())
	}
	switch pub := pubRSA.(type){
	case *rsa.PublicKey:
		recv.pubRSA = pub
	default:
		return errors.New(reflect.TypeOf(recv).String() + ": " + "error reading public key")
	}
	return nil
}

func (recv *EncrypterImpl)setElGamal(keyData []byte)error{
	publicKeyElGamal := &elgamal.PublicKey{}

	block, keyData := pem.Decode(keyData)
	publicKeyElGamal.G = new(big.Int).SetBytes(block.Bytes)

	block, keyData = pem.Decode(keyData)
	publicKeyElGamal.P = new(big.Int).SetBytes(block.Bytes)

	block, _ = pem.Decode(keyData)
	publicKeyElGamal.Y = new(big.Int).SetBytes(block.Bytes)

	recv.pubElGamal = publicKeyElGamal

	return nil
}

func (recv *EncrypterImpl)Encrypt(encType crypt.TYPE, msg []byte)([]byte, error){
	switch encType {
	case crypt.PLAIN:
		return msg, nil
	case crypt.RSA:
		return crypt.EncryptRSA(recv.pubRSA, msg)
	case crypt.ELGAMAL:
		return crypt.EncryptElGamal(recv.pubElGamal, msg)
	}

	return msg, nil
}

func (enc *EncrypterImpl)GetPublicKeyRSA() *rsa.PublicKey{
	return enc.pubRSA;
}

func (enc *EncrypterImpl)GetPublicKeyElGamal() *elgamal.PublicKey{
	return enc.pubElGamal;
}
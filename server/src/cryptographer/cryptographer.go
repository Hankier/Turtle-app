package cryptographer

import (
	"crypto/rsa"
	"log"
	"crypto/rand"
	"golang.org/x/crypto/openpgp/elgamal"
	"math/big"
	"io"
	"errors"
)

type TYPE byte

const (
	PLAIN TYPE = iota
	RSA
	ELGAMAL
)

//https://tools.ietf.org/html/rfc5114#section-2.1
var p = fromHex("87A8E61DB4B6663CFFBBD19C651959998CEEF608660DD0F25D2CEED4435E3B00E00DF8F1D61957D4FAF7DF4561B2AA3016C3D91134096FAA3BF4296D830E9A7C209E0C6497517ABD5A8A9D306BCF67ED91F9E6725B4758C022E0B1EF4275BF7B6C5BFC11D45F9088B941F54EB1E59BB8BC39A0BF12307F5C4FDB70C581B23F76B63ACAE1CAA6B7902D52526735488A0EF13C6D9A51BFA4AB3AD8347796524D8EF6A167B5A41825D967E144E5140564251CCACB83E6B486F6B3CA3F7971506026C0B857F689962856DED4010ABD0BE621C3A3960A54E710C375F26375D7014103A4B54330C198AF126116D2276E11715F693877FAD7EF09CADB094AE91E1A1597")
var pLen = len(p.Bytes())
var g = fromHex("3FB32C9B73134D0B2E77506660EDBD484CA7B18F21EF205407F4793A1A0BA12510DBC15077BE463FFF4FED4AAC0BB555BE3A6C1B0C6B47B1BC3773BF7E8C6F62901228F8C28CBB18A55AE31341000A650196F931C77A57F2DDF463E5E9EC144B777DE62AAAB8A8628AC376D282D6ED3864E67982428EBC831D14348F6F2F9193B5045AF2767164E1DFC967C1FB3F2E55A4BD1BFFE83B9C80D052B985D182EA0ADB2A3B7313D3FE14C8484B1E052588B9B7D2BBD2DF016199ECD06E1557CD0915B3353BBB64E0EC377FD028370DF92B52C7891428CDC67EB6184B523D1DB246C32F63078490F00EF8D647D148D47954515E2327CFEF98C582664B4C0F6CC41659")

func fromHex(hex string) *big.Int {
	n, ok := new(big.Int).SetString(hex, 16)
	if !ok {
		panic("failed to parse hex number")
	}
	return n
}

type Cryptographer interface{
	Encrypt(encType TYPE, bytes []byte)[]byte
	Decrypt(encType TYPE, bytes []byte)[]byte
}

func GenerateRSA() *rsa.PrivateKey {
	privateKeyRSA, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatal(err)
	}

	privateKeyRSA.Precompute()

	if err = privateKeyRSA.Validate(); err != nil {
		log.Fatal(err)
	}

	return privateKeyRSA
}

func GenerateElGamal() *elgamal.PrivateKey{
	x, _ := rand.Int(rand.Reader, p)
	privateKeyElGamal := &elgamal.PrivateKey{
		PublicKey: elgamal.PublicKey{
			G: g,
			P: p,
		},
		X: x,
	}
	privateKeyElGamal.Y = new(big.Int).Exp(privateKeyElGamal.G, privateKeyElGamal.X, privateKeyElGamal.P)

	return privateKeyElGamal
}

func DecryptRSA(privateKey *rsa.PrivateKey, msg []byte)[]byte{
	decryptedText, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, msg);
	if  err != nil {
		log.Fatal(err)
	}
	return decryptedText
}

func DecryptElGamal(privateKey *elgamal.PrivateKey, msg []byte) ([]byte, error) {
	if len(msg) != pLen{
		return nil, errors.New("bad message length")
	}
	c1 := new(big.Int).SetBytes(msg[0:pLen / 2])
	c2 := new(big.Int).SetBytes(msg[pLen / 2:pLen])
	return elgamal.Decrypt(privateKey, c1, c2)
}

func DecryptPlain(msg []byte) []byte {
	return msg
}

func EncryptRSA(publicKey *rsa.PublicKey, msg []byte)[]byte{
	encryptedText, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, msg);
	if  err != nil {
		log.Fatal(err)
	}
	return encryptedText
}

func EncryptElGamal(publicKey *elgamal.PublicKey, msg []byte) ([]byte, error) {
	c1, c2, err := elgamal.Encrypt(rand.Reader, publicKey, msg)
	if err != nil{
		return nil, err
	}
	encmsg := c1.Bytes()
	encmsg = append(encmsg, c2.Bytes()...)
	return encmsg, nil
}

func EncryptPlain(msg []byte) []byte {
	return msg
}

func nonZeroRandomBytes(s []byte, rand io.Reader) (err error) {
	_, err = io.ReadFull(rand, s)
	if err != nil {
		return
	}

	for i := 0; i < len(s); i++ {
		for s[i] == 0 {
			_, err = io.ReadFull(rand, s[i:i+1])
			if err != nil {
				return
			}
		}
	}

	return
}


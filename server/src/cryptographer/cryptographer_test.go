package cryptographer

import (
	"testing"
	"math/big"
	"golang.org/x/crypto/openpgp/elgamal"
	"crypto/rand"
	"bytes"
)

func TestElGamal(t *testing.T) {
	x, _ := rand.Int(rand.Reader, p)

	privKey := &elgamal.PrivateKey{
		PublicKey: elgamal.PublicKey{
			G: g,
			P: p,
		},
		X: x,
	}

	privKey.Y = new(big.Int).Exp(privKey.G, privKey.X, privKey.P)

	message := make([]byte, 245)
	enc, err := EncryptElGamal(&privKey.PublicKey, message)
	if err != nil {
		t.Error(err)
	}
	message2, err := DecryptElGamal(privKey, enc)
	if err != nil {
		t.Error(err)
	}
	if !bytes.Equal(message2, message){
		t.Error("Decryption failed, got: ", message2, " expected: ", message)
	}
}

func TestElGamalTooLongMessage(t *testing.T) {
	x, _ := rand.Int(rand.Reader, p)

	privKey := &elgamal.PrivateKey{
		PublicKey: elgamal.PublicKey{
			G: g,
			P: p,
		},
		X: x,
	}

	privKey.Y = new(big.Int).Exp(privKey.G, privKey.X, privKey.P)

	message := make([]byte, 246)
	_, err := EncryptElGamal(&privKey.PublicKey, message)
	if err == nil {
		t.Error("Should return error")
	}
}
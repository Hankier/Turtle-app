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

func TestElGamalLoadSave(t *testing.T){
	x, _ := rand.Int(rand.Reader, p)

	privKey := &elgamal.PrivateKey{
		PublicKey: elgamal.PublicKey{
			G: g,
			P: p,
		},
		X: x,
	}

	privKey.Y = new(big.Int).Exp(privKey.G, privKey.X, privKey.P)

	SaveElGamal(privKey, "testFile")
	privKey2, _ := LoadElGamal("testFile")

	if privKey.G.Cmp(privKey2.G) != 0 || privKey.P.Cmp(privKey2.P) != 0 || privKey.X.Cmp(privKey2.X) != 0 || privKey.Y.Cmp(privKey2.Y) != 0{
		t.Error("keys do not match!")
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
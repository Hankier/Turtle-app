package decrypter

const (
	PLAIN = iota
	RSA
	ELGAMAL
)

type Decrypter interface{
	Decrypt(int, []byte)([]byte)
}

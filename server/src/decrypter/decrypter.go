package decrypter

type TYPE int

const (
	PLAIN TYPE = iota
	RSA
	ELGAMAL
)

type Decrypter interface{
	Decrypt(TYPE, []byte)([]byte)
}

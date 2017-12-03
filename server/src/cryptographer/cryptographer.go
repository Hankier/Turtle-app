package cryptographer

type TYPE byte

const (
	PLAIN TYPE = iota
	RSA
	ELGAMAL
)

type Decrypter interface{
	Decrypt(TYPE, []byte)([]byte)
}

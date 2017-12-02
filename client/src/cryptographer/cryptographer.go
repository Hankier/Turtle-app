package cryptographer

type TYPE byte

const (
	PLAIN TYPE = iota
	RSA
	ELGAMAL
)

type Cryptographer interface{
	encrypt([]byte)[]byte
	decrypt([]byte)[]byte
}

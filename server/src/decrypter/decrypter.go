package decrypter

type Decrypter interface{
	Decrypt([]byte)([]byte)
}

package decrypterInterface

type Decrypter interface{
	Decrypt([]byte)([]byte)
}

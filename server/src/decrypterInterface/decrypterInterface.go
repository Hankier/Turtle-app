package decrypterInterface

type Decrypter interface{
	decrypt([]byte)([]byte)
}

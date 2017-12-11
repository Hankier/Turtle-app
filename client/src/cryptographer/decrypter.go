package cryptographer

type Decrypter interface{
	Decrypt(encType TYPE, bytes []byte) ([]byte, error)
}

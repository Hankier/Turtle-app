package crypt

type Decrypter interface{
	Decrypt(encType TYPE, bytes []byte) ([]byte, error)
}

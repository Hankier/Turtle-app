package cryptographer

type Encrypter interface{
	Encrypt(encType TYPE, bytes []byte) ([]byte, error)
}

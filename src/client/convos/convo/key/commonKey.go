package key

type TYPE byte

const (
	PLAIN TYPE = iota
	SYMMETRIC
)

type CommonKey interface{
	Decrypt(encType TYPE, bytes []byte)([]byte, error)
	decryptSymmetric(bytes []byte)([]byte, error)
	SetCommonKeyData(part int, bytes []byte)
	GetCommonKeyData(part int)([]byte)
}
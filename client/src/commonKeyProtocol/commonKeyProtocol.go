package commonKeyProtocol

type TYPE byte

const (
	PLAIN TYPE = iota
	SYMMETRIC
)

type CommonKeyProtocol interface{
	Decrypt(bytes []byte)([]byte)
	decryptSymmetric(bytes []byte)([]byte, error)
	SetCommonKeyData(part int, bytes []byte)
	GetCommonKeyData(part int)([]byte)
}
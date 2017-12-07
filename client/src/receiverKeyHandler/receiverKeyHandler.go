package receiverKeyHandler

import "cryptographer"

type ReceiverKeyHandler interface{
	SetKey(ktype cryptographer.TYPE, bytes []byte)
	Encrypt(ktype cryptographer.TYPE, bytes []byte)([]byte, error)
}
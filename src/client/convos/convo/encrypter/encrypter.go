package encrypter

import "crypt"

type Encrypter interface{
	crypt.Encrypter
	SetKey(p crypt.TYPE, key []byte)error
}

package key

import (
	"github.com/Nik-U/pbc"
)
type CommonKeyImpl struct{
	pairing                  *pbc.Pairing
	gen                      *pbc.Element //g
	sessionGen               *pbc.Element //g dash
	secretKey                *pbc.Element //a, b
	publicKey                *pbc.Element //A / B
	otherPublicKey           *pbc.Element //A / B
	sessionGenSecretKey      *pbc.Element //x, y
	sessionGenPublicKey      *pbc.Element //X / Y
	otherSessionGenPublicKey *pbc.Element //X / Y
	sessionMultiplier        *pbc.Element //ca / cb
	otherSessionMultiplier   *pbc.Element //ca / cb
	sessionPublicKey         *pbc.Element //Sa / Sb
	otherSessionPublicKey    *pbc.Element //Sa / Sb
	sessionCommonKey         *pbc.Element //K
}

func New()(*CommonKeyImpl){
	//TODO
	ckpi := new(CommonKeyImpl)

	return ckpi;
}

func (ckpi *CommonKeyImpl)Decrypt(encType TYPE, bytes []byte)([]byte, error){
	switch encType {
	case PLAIN:
		return bytes, nil
	case SYMMETRIC:
		return ckpi.decryptSymmetric(bytes)
	}
	return bytes, nil
}

func (ckpi *CommonKeyImpl)decryptSymmetric(bytes []byte)([]byte, error){
	//TODO
	return nil, nil
}
func (ckpi *CommonKeyImpl)SetCommonKeyData(part int, bytes []byte){
	//TODO
}
func (ckpi *CommonKeyImpl)GetCommonKeyData(part int)([]byte){
	//TODO
	return nil
}
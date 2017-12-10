package commonKeyProtocol

import (
	"github.com/Nik-U/pbc"
)
type CommonKeyProtocolImpl struct{
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

func NewCommonKeyProtocolImpl()(*CommonKeyProtocolImpl){
	//TODO
	ckpi := new(CommonKeyProtocolImpl)

	return ckpi;
}

func (ckpi *CommonKeyProtocolImpl)Decrypt(encType TYPE, bytes []byte)([]byte, error){
	switch encType {
	case PLAIN:
		return bytes, nil
	case SYMMETRIC:
		return ckpi.decryptSymmetric(bytes)
	}
	return bytes, nil
}

func (ckpi *CommonKeyProtocolImpl)decryptSymmetric(bytes []byte)([]byte, error){
	//TODO
	return nil, nil
}
func (ckpi *CommonKeyProtocolImpl)SetCommonKeyData(part int, bytes []byte){
	//TODO
}
func (ckpi *CommonKeyProtocolImpl)GetCommonKeyData(part int)([]byte){
	//TODO
	return nil
}
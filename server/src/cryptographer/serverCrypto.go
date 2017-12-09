package cryptographer

import (
    "crypto/rsa"
    "golang.org/x/crypto/openpgp/elgamal"
)

type ServerCrypto struct{
    privateKeyRSA  *rsa.PrivateKey
    publicKeyRSA   *rsa.PublicKey
    privateKeyElGamal  *elgamal.PrivateKey
    publicKeyElGamal   *elgamal.PublicKey
}


func NewServerCrypto()(*ServerCrypto){
    srv := new(ServerCrypto)

    var privateKeyRSA *rsa.PrivateKey

	privateKeyRSA, err := LoadRSA("privateKeyRSA")
    if err != nil{
        privateKeyRSA = GenerateRSA()
        SaveRSA(privateKeyRSA, "privateKeyRSA")
    }
	srv.privateKeyRSA = privateKeyRSA
	srv.publicKeyRSA = &privateKeyRSA.PublicKey


	var privateKeyElGamal *elgamal.PrivateKey

	privateKeyElGamal, err = LoadElGamal("privateKeyElGamal")
	if err != nil{
		privateKeyElGamal = GenerateElGamal()
		SaveElGamal(privateKeyElGamal, "privateKeyElGamal")
	}
	srv.privateKeyElGamal = privateKeyElGamal
	srv.publicKeyElGamal = &privateKeyElGamal.PublicKey

    return srv
}

func (srv *ServerCrypto)Decrypt(encType TYPE, bytes []byte)[]byte{
    switch encType {
    case PLAIN:
        return DecryptPlain(bytes)
    case RSA:
        return DecryptRSA(srv.privateKeyRSA, bytes)
    case ELGAMAL:
        //return DecryptElGamal(bytes)
    }
    return bytes
}

func (srv *ServerCrypto)Encrypt(encType TYPE, bytes []byte)[]byte{
	switch encType {
	case PLAIN:
		return EncryptPlain(bytes)
	case RSA:
		return EncryptRSA(srv.publicKeyRSA, bytes)
	case ELGAMAL:
		//return EncryptElGamal(bytes)
	}
	return bytes
}

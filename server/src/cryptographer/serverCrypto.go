package cryptographer

import (
    "log"
    "io/ioutil"
    "crypto/rsa"
    "crypto/rand"
    "encoding/pem"
    "crypto/x509"
)

type ServerCrypto struct{
    privateKey  *rsa.PrivateKey
    publicKey   *rsa.PublicKey
}


func NewServerCrypto()(*ServerCrypto){
    srv := new(ServerCrypto)
    if !srv.loadKey(){
        srv.generateKey()
    }
    return srv
}

func (srv *ServerCrypto)generateKey(){
    privateKey, err := rsa.GenerateKey(rand.Reader, 1024)
    if err != nil {
        log.Fatal(err)
    }

    privateKey.Precompute()

    if err = privateKey.Validate(); err != nil {
        log.Fatal(err)
    }

    pemdata := pem.EncodeToMemory(
    &pem.Block{
        Type: "RSA PRIVATE KEY",
        Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
    },
    )
    err = ioutil.WriteFile("_privateKey",pemdata,0644)
}



func (srv *ServerCrypto)loadKey() bool{
    msg, err := ioutil.ReadFile("_privateKey")
    if err != nil {
        log.Print("Error reading private key.")
        return false
    }
    block, _ := pem.Decode(msg)

    priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
    if err != nil {
        log.Print("Failed to parse private key: " + err.Error())
        return false
    }
    priv.Precompute()

    if err = priv.Validate(); err != nil {
        log.Fatal(err)
        return false
    }
    srv.privateKey = priv
    srv.publicKey = &priv.PublicKey

    return true
}

func (srv *ServerCrypto)Decrypt(encType TYPE, bytes []byte)[]byte{
    switch encType {
    case PLAIN:
        return DecryptPlain(bytes)
    case RSA:
        return DecryptRSA(srv.privateKey, bytes)
    case ELGAMAL:
        return DecryptElGamal(bytes)
    }
    return bytes
}

func (srv *ServerCrypto)Encrypt(encType TYPE, bytes []byte)[]byte{
	switch encType {
	case PLAIN:
		return EncryptPlain(bytes)
	case RSA:
		return EncryptRSA(srv.publicKey, bytes)
	case ELGAMAL:
		return EncryptElGamal(bytes)
	}
	return bytes
}

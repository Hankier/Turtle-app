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

func (srv *ServerCrypto)decryptRSA(msg []byte)[]byte{
    decryptedText, err := rsa.DecryptPKCS1v15(rand.Reader, srv.privateKey, msg);
    if  err != nil {
        log.Fatal(err)
    }
    return decryptedText
}

func (srv *ServerCrypto)encryptRSA(msg []byte)[]byte{
	encryptedText, err := rsa.EncryptPKCS1v15(rand.Reader, srv.publicKey, msg)
    if err != nil {
        log.Fatal(err)
    }
    return encryptedText
}

func (srv *ServerCrypto)Decrypt(enctype TYPE, bytes []byte)[]byte{
    switch enctype {
    case PLAIN:
        return srv.decryptPlain(bytes)
    case RSA:
        return srv.decryptRSA(bytes)
    case ELGAMAL:
        return srv.decryptElGamal(bytes)
    }
    return bytes
}

func (srv *ServerCrypto) decryptElGamal(bytes []byte) []byte {
    return bytes
    //TODO ELGAMAL!!!
}

func (srv *ServerCrypto) decryptPlain(bytes []byte) []byte {
    return bytes
}

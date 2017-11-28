package decrypter

import (
    "log"
    "io/ioutil"
    "hash"
    "crypto/md5"
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
    generateKey()
    loadKey()
    return srv
}

func (srv *ServerCrypto)generateKey()bool{
    if privateKey, err = rsa.GenerateKey(rand.Reader, 1024); err != nil {
        log.Fatal(err)
    }

    privateKey.Precompute()

    if err = privateKey.Validate(); err != nil {
        log.Fatal(err)
    }

    pemdata := pem.EncodeToMemory(
    &pem.Block{
        Type: "RSA PRIVATE KEY",
        Bytes: x509.MarshalPKCS1PrivateKey(key),
    },
    )
    err = ioutil.WriteFile("_privateKey",pemdata,0644)
}



func (srv *ServerCrypto)loadKey()bool{
    var bytes = make([]byte, 1024)
    bytes, _ = ioutil.ReadFile("_privateKey")
    if err != nil {
        log.Print("Error reading private key.")
        return false
    }
    block, _ := pem.Decode([]byte(read_bs))

    priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
    if err != nil {
        log.Println("Failed to parse private key: %s", err)
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

func (srv *ServerCrypto)DecryptRSA(bytes []byte)[]byte{
    var err error
    var md5_hash hash.Hash
    var label []byte
    if decryptedText, err = rsa.DecryptOAEP(md5_hash, rand.Reader, srv.privateKey, bytes, label); err != nil {
        log.Fatal(err)
    }
    return decryptedText
}

func (srv *ServerCrypto)EncrytRSA(bytes []byte)[]byte{
    var err error
    var md5_hash hash.Hash
    var label []byte
    if encryptedText, err = rsa.EncrytOAEP(md5_hash, rand.Reader, srv.publicKey, bytes, label); err != nil {
        log.Fatal(err)
    }
    return encryptedText
}

func (srv *ServerCrypto)Decrypt(bytes []byte)[]byte{

    return bytes
}

func main() {
    serv := new(ServerCrypto)


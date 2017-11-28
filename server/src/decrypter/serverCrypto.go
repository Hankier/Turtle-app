package decrypter

import (
    "log"
    "io/ioutil"
)

//type ServerCrypto struct{
//    privateKey  []byte
//    publicKey   []byte
//}

type ServerCrypto struct{
}

func NewServerCrypto()(*ServerCrypto){
    srv := new(ServerCrypto)
    //srv.loadKey()

    return srv
}

func (srv *ServerCrypto)generateKey()bool{


func (srv *ServerCrypto)loadKey()bool{
    var bytes = make([]byte, 1024)
    bytes, _ = ioutil.ReadFile("_privateKey")
    if err != nil {
        log.Print("Error reading private key.")
        return false
    }
    srv.privateKey.SetBytes(bytes[:])
    bytes, _ = ioutil.ReadFile("_publicKey")
    if err != nil {
        log.Print("Error reading public key.")
        return false
    }
    srv.publicKey.SetBytes(bytes[:])
    return true
}

func (srv *ServerCrypto)Decrypt(bytes []byte)[]byte{

    return bytes
}

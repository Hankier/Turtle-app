package entry

import (
	"crypto/rsa"
	"golang.org/x/crypto/openpgp/elgamal"
	"crypt"
	"srvlist/encrypter"
)
type Entry struct{
	Name      string
	Ipport    string
	Encrypter crypt.Encrypter
}

func New(name string, ip_port string, pubRSA *rsa.PublicKey, pubElGamal *elgamal.PublicKey)(*Entry){
	srvEntry := new(Entry)
	srvEntry.Ipport = ip_port
	srvEntry.Name = name
	srvEntry.Encrypter = encrypter.New(pubRSA, pubElGamal)
	return srvEntry
}


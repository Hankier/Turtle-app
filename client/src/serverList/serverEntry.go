package serverList

import (
	"crypto/rsa"
	"golang.org/x/crypto/openpgp/elgamal"
	)
type serverEntry struct{
	Name string
	Ip_port string
	PublicKeyRSA *rsa.PublicKey
	PublicKeyElGamal *elgamal.PublicKey
}

func NewServerEntry(name string, ip_port string, pkRSA *rsa.PublicKey, pkElGamal *elgamal.PublicKey)(*serverEntry){
	srvEntry := new(serverEntry)
	srvEntry.Ip_port = ip_port
	srvEntry.Name = name
	srvEntry.PublicKeyElGamal = pkElGamal
	srvEntry.PublicKeyRSA = pkRSA
	return srvEntry
}

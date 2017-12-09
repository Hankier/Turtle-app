package serverList

import (
	"sync"
	"crypto/rsa"
	"golang.org/x/crypto/openpgp/elgamal"
	"math/big"
	"crypto/rand"
)

type ServerList struct{
	serverListMutex sync.Mutex
	serverList map[string]*serverEntry
	daList []string
}

func NewServerList()(*ServerList)  {
	sli := new(ServerList)

	//TODO debug remove hardcoded serverEntry info

	sli.serverList = make(map[string]*serverEntry)

	sli.serverList["00000000"] = NewServerEntry("00000000", "127.0.0.1:8081", nil, nil)
	sli.serverList["00000001"] = NewServerEntry("00000001", "127.0.0.1:8083", nil, nil)
	sli.serverList["00000002"] = NewServerEntry("00000002", "127.0.0.1:8085", nil, nil)

	return sli
}

func (sli *ServerList)GetServerIpPort(name string)(string){
	return sli.serverList[name].Ip_port
}

func (sli *ServerList)GetPublicKeyRSA(name string)(*rsa.PublicKey){
	return sli.serverList[name].PublicKeyRSA
}

func (sli *ServerList)GetPublicKeyElGamal(name string)(*elgamal.PublicKey){
	return sli.serverList[name].PublicKeyElGamal
}

func (sli *ServerList)GetRandomPath(length int)([]string){
	path := make([]string, length)

	keys := make([]string, 0, len(sli.serverList))
	for k := range sli.serverList {
		keys = append(keys, k)
	}

	var key string
	var rnd *big.Int


	for i := 0; i < length; i++{
		rnd, _ = rand.Int(rand.Reader, big.NewInt(int64(len(keys))))
		key = keys[rnd.Int64()]
		path[i] = sli.serverList[key].Name
		for i > 0 && path[i] == path[i-1]{
			rnd, _ = rand.Int(rand.Reader, big.NewInt(int64(len(keys))))
			key = keys[rnd.Int64()]
			path[i] = sli.serverList[key].Name
		}
	}


	return path
}

func (sli *ServerList)RefreshList(){
	//TODO
 }
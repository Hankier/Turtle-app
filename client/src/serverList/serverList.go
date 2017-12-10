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
}

func NewServerList()(*ServerList)  {
	sli := new(ServerList)

	//TODO debug remove hardcoded serverEntry info

	sli.serverList = make(map[string]*serverEntry)

	sli.serverList["00000000"] = NewServerEntry("00000000", "127.0.0.1:8080", nil, nil)
	sli.serverList["00000001"] = NewServerEntry("00000001", "127.0.0.1:8082", nil, nil)
	sli.serverList["00000002"] = NewServerEntry("00000002", "127.0.0.1:8084", nil, nil)

	return sli
}

func (sli *ServerList)GetServerIpPort(name string)(string){
	sli.serverListMutex.Lock()
	ret := sli.serverList[name].Ip_port
	sli.serverListMutex.Unlock()
	return ret
}

func (sli *ServerList)GetPublicKeyRSA(name string)(*rsa.PublicKey){
	sli.serverListMutex.Lock()
	ret := sli.serverList[name].PublicKeyRSA
	sli.serverListMutex.Unlock()
	return ret
}

func (sli *ServerList)GetPublicKeyElGamal(name string)(*elgamal.PublicKey){
	sli.serverListMutex.Lock()
	ret := sli.serverList[name].PublicKeyElGamal
	sli.serverListMutex.Unlock()
	return ret
}

func (sli *ServerList)GetRandomPath(length int)([]string){
	path := make([]string, length)

	names := sli.GetServerList()

	var rnd *big.Int

	for i := 0; i < length; i++{
		rnd, _ = rand.Int(rand.Reader, big.NewInt(int64(len(names))))
		path[i] = names[rnd.Int64()]
		for i > 0 && path[i] == path[i-1]{
			rnd, _ = rand.Int(rand.Reader, big.NewInt(int64(len(names))))
			path[i] = names[rnd.Int64()]
		}
	}

	return path
}

func (sli *ServerList)GetServerList()[]string{
	names := make([]string, 0, len(sli.serverList))
	sli.serverListMutex.Lock()
	for k := range sli.serverList {
		names = append(names, k)
	}
	sli.serverListMutex.Unlock()
	return names
}

func (sli *ServerList)RefreshList(){
	//TODO
 }
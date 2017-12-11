package serverList

import (
	"sync"
	"crypto/rsa"
	"golang.org/x/crypto/openpgp/elgamal"
	"math/big"
	"crypto/rand"
	"errors"
)

type ServerList struct{
	serverListMutex sync.Mutex
	serverList map[string]*serverEntry
}

func NewServerList()(*ServerList)  {
	sli := new(ServerList)

	//TODO debug remove hardcoded serverEntry info

	sli.serverList = make(map[string]*serverEntry)

	sli.serverList["00000000"] = NewServerEntry("00000000", "127.0.0.1:8080", nil, nil, nil)
	sli.serverList["00000001"] = NewServerEntry("00000001", "127.0.0.1:8082", nil, nil, nil)
	sli.serverList["00000002"] = NewServerEntry("00000002", "127.0.0.1:8084", nil, nil, nil)

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

func (sli *ServerList)GetRandomPath(length int)([]string, error){
	if length < 1{
		if length < 0{
			return nil, errors.New("Invalid path length")
		}
		return make([]string, 0), nil
	}

	path := make([]string, length)

	names := sli.GetServerList()


	serversLen := len(names)

	if serversLen < 2{
		return nil, errors.New("Too few servers to create a path");
	}

	var rnd *big.Int
	var err error

	for i := 1; i < length; i++{
		rnd, err = rand.Int(rand.Reader, big.NewInt(int64(serversLen)))
		if err != nil {	return nil, err	}
		path[i] = names[rnd.Int64()]
		for path[i] == path[i-1]{
			rnd, err = rand.Int(rand.Reader, big.NewInt(int64(serversLen)))
			if err != nil {	return nil, err	}
			path[i] = names[rnd.Int64()]
		}
	}

	return path, nil
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
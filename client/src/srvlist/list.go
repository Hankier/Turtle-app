package srvlist

import (
	"sync"
	"math/big"
	"crypto/rand"
	"errors"
	"crypt"
	"srvlist/entry"
)

type ServerList struct{
	listmutex sync.Mutex
	list      map[string]*entry.Entry
}

func New()(*ServerList)  {
	sli := new(ServerList)

	//TODO debug remove hardcoded serverEntry info

	sli.list = make(map[string]*entry.Entry)

	sli.list["00000000"] = entry.New("00000000", "127.0.0.1:8080", nil, nil)
	sli.list["00000001"] = entry.New("00000001", "127.0.0.1:8082", nil, nil)
	sli.list["00000002"] = entry.New("00000002", "127.0.0.1:8084", nil, nil)

	return sli
}

func (sli *ServerList)GetServerIpPort(name string)(string, error){

	sli.listmutex.Lock()
	ret, ok := sli.list[name];
	sli.listmutex.Unlock()

	if  ok{
		return ret.Ipport, nil
	}

	return "", errors.New("no such server on the list")
}

func (sli *ServerList)GetEncrypter(name string)(crypt.Encrypter, error){
	sli.listmutex.Lock()
	entry, ok := sli.list[name]
	sli.listmutex.Unlock()

	if ok{
		return entry.Encrypter, nil
	}
	return nil, errors.New("no such server on the list")
}

func (sli *ServerList)GetRandomPath(length int)([]string, error){
	if length < 1{
		if length < 0{
			return nil, errors.New("invalid path length")
		}
		return make([]string, 0), nil
	}

	path := make([]string, length)

	//no need for mutex, GetServerList is thread-safe
	names := sli.GetServerList()

	serversLen := len(names)

	if serversLen < 2{
		return nil, errors.New("too few servers to create a path");
	}

	var rnd *big.Int
	var err error

	for i := 0; i < length; i++{

		rnd, err = rand.Int(rand.Reader, big.NewInt(int64(serversLen)))
		if err != nil {	return nil, err	}

		path[i] = names[rnd.Int64()]

		for i > 0 && path[i] == path[i-1]{

			rnd, err = rand.Int(rand.Reader, big.NewInt(int64(serversLen)))
			if err != nil {	return nil, err	}

			path[i] = names[rnd.Int64()]
		}
	}

	return path, nil
}

func (sli *ServerList)GetServerList()[]string{
	names := make([]string, 0, len(sli.list))
	sli.listmutex.Lock()
	for k := range sli.list {
		names = append(names, k)
	}
	sli.listmutex.Unlock()
	return names
}

func (sli *ServerList)RefreshList(){
	//TODO
 }
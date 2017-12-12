package srvlist

import (
	"sync"
	"srvlist/entry"
	"errors"
)


type ServerList struct{
	listmutex sync.Mutex
	list      map[string]*entry.Entry
}

func New()(*ServerList)  {
	sli := new(ServerList)

	//TODO debug remove hardcoded serverEntry info

	sli.list = make(map[string]*entry.Entry)

	sli.list["00000000"] = entry.New("00000000", "127.0.0.1:8080", nil)
	sli.list["00000001"] = entry.New("00000001", "127.0.0.1:8082", nil)
	sli.list["00000002"] = entry.New("00000002", "127.0.0.1:8084", nil)

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
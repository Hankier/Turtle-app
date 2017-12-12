package serverList

import (
	"sync"
	"entry"
	"strconv"
)

type ServerList struct{
	listmutex sync.Mutex
	list      map[string]*entry.Entry
}


func (sli *ServerList)GetServerIpPort(name string)(string){
	sli.listmutex.Lock()
	ret := sli.list[name].Ipport
	sli.listmutex.Unlock()
	return ret
}

func (sli *ServerList)GetEncrypter(name string)(string){
	sli.listmutex.Lock()
	ret := sli.list[name].Encrypter
	sli.listmutex.Unlock()
	return ret
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

func (sli *ServerList)AddServerToList(ip_port string, pk []byte)(*ServerList){
	next_free := int(len(sli.list))
	t := strconv.Itoa(next_free)
  //TODO change next_name setting 
	next_name := "0000000" + t

	sli.listmutex.Lock()
	sli.list[next_name] = entry.New(next_name, ip_port, pk)
	sli.listmutex.Unlock()

	return sli
}

func New()(*ServerList)  {
	sli := new(ServerList)

	//TODO debug remove hardcoded serverEntry info

	sli.list = make(map[string]*entry.Entry)

	sli.list["00000000"] = entry.New("00000000", "127.0.0.1:8080", []byte("aa"))
	sli.list["00000001"] = entry.New("00000001", "127.0.0.1:8082", []byte("bb"))
	sli.list["00000002"] = entry.New("00000002", "127.0.0.1:8084", []byte("cc"))

	return sli
}

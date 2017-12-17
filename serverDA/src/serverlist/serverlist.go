package serverlist

import (
	"sync"
	"../entry"
	"strconv"
	"strings"
)

type ServerList struct{
	listmutex sync.Mutex
	list      map[string]*entry.Entry
}


func (sli *ServerList)GetServerIpPort(name string)(string){
	sli.listmutex.Lock()
	ret := sli.list[name].Ip_port
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
	next_free_str := strconv.Itoa(next_free)
	zeros_number := 8 - len(next_free_str)
	next_name := strings.Repeat("0", zeros_number) + next_free_str

	sli.listmutex.Lock()
	sli.list[next_name] = entry.NewEntry(next_name, ip_port, pk)
	sli.listmutex.Unlock()

	return sli
}


func (sli *ServerList)RemoveServerFromList(name string)(*ServerList){

	delete(sli.list, name)

	return sli
}

func NewList()(*ServerList)  {
	sli := new(ServerList)


	sli.list = make(map[string]*entry.Entry)

	return sli
}

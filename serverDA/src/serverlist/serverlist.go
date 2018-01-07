package serverlist

import (
	"sync"
	"entry"
	"strconv"
	"strings"
	"errors"
	"encoding/json"
	"io/ioutil"
	"log"
)

type ServerList struct{
	listmutex sync.Mutex
	list      map[string]*entry.Entry
}


func (sli *ServerList)GetServerIpPort(name string)(string, error){

	sli.listmutex.Lock()
	ret, ok := sli.list[name];
	sli.listmutex.Unlock()

	if  ok{
		return ret.Ip_port, nil
	}

	return "", errors.New("no such server on the list")
}

func (sli *ServerList)GetServerKey(name string)(string){
	sli.listmutex.Lock()
	ret := sli.list[name].PublicKey
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


func (sli *ServerList)GetServerListJSON()[]byte{
	json_list, _ := json.Marshal(sli.list)

	ioutil.WriteFile("list_of_servers.json",json_list,0644)

	return json_list
}

func (sli *ServerList)AddServerToList(ip_port string, pk string)string{
	next_free := int(len(sli.list))
	next_free_str := strconv.Itoa(next_free)
	zeros_number := 8 - len(next_free_str)
	next_name := strings.Repeat("0", zeros_number) + next_free_str

	sli.listmutex.Lock()
	sli.list[next_name] = entry.NewEntry(next_name, ip_port, pk)
	sli.listmutex.Unlock()

	return next_name
}


func (sli *ServerList)RemoveServerFromList(name string)(*ServerList){

	delete(sli.list, name)

	return sli
}

func (sli *ServerList)SaveListToFile(filename string) (err error) {

	json_list, _ := json.Marshal(sli.list)

	ioutil.WriteFile(filename,json_list,0644)
	if err == nil{
		log.Print("List saved to: ", filename)
	}
	return err
}

//func GetListFromFile(filename string) (sli *ServerList,err error) {

	//TODO get list from file

//}

func NewList()(*ServerList)  {
	sli := new(ServerList)


	sli.list = make(map[string]*entry.Entry)

	return sli
}

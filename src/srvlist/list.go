package srvlist

import (
	"sync"
	"math/big"
	"crypto/rand"
	"errors"
	"crypt"
	"srvlist/entry"
	"strconv"
	"log"
	"strings"
	"crypto/rsa"
	"golang.org/x/crypto/openpgp/elgamal"
	"io/ioutil"
	"reflect"
)

//ServerList class handling operations on server entries thread-safely
type ServerList struct{
	listmutex sync.Mutex				//mutex to keep on-list operations thread-safe
	list      map[string]*entry.Entry	//a map of server entries with names as keys
}


//ServerList constructor
//Creates a map of server entries with names as keys
func New()(*ServerList)  {
	sli := new(ServerList)

	sli.list = make(map[string]*entry.Entry)

	return sli
}

func (sli *ServerList)SetList(list map[string]*entry.Entry){
	sli.list = list;
}

//GetServerIpPort returns a string with with format ip:port of server of given name or an error if there is no server of given name
func (sli *ServerList)GetServerIpPort(name string)(string, error){

	sli.listmutex.Lock()
	ret, ok := sli.list[name];
	sli.listmutex.Unlock()

	if  ok{
		return ret.Ipport, nil
	}

	return "", errors.New(reflect.TypeOf(sli).String() + ": no such server on the list")
}


//GetEncrypter returns an encrypter object of server of a given name or an error if there is no server of given name
func (sli *ServerList)GetEncrypter(name string)(crypt.Encrypter, error){
	sli.listmutex.Lock()
	entr, ok := sli.list[name]
	sli.listmutex.Unlock()

	if ok{
		return entr.Encrypter, nil
	}
	return nil, errors.New(reflect.TypeOf(sli).String() +": no such server on the list")
}

//Generates a cryptographically secure random path and returnes it as a slice of strings representing names of consecutive nodes(servers)
//Returns mentioned slice and nil if all went well
//Returns nil and error in the following cases:
// -path length is smaller than 0
// -there are less than two servers to generate path from (if path length is longer than 1), because two consecutive servers cannot be same
// -occurred an error creating crypto/rand number
func (sli *ServerList)GetRandomPath(length int)([]string, error){
	if length < 1{
		if length < 0{
			return nil, errors.New(reflect.TypeOf(sli).String() + ": invalid path length")
		}
		return make([]string, 0), nil
	}
	//no need for mutex, GetServerList is thread-safe
	names := sli.GetServerList()

	serversLen := len(names)

	if serversLen < 2 && length > 1{
		return nil, errors.New(reflect.TypeOf(sli).String() + ":too few servers to create a path");
	}

	path := make([]*big.Int, length)

	var rnd *big.Int
	var err error

	srvListLen := big.NewInt(int64(serversLen))

	rnd, err = rand.Int(rand.Reader, srvListLen)
	if err != nil {	return nil, err	}

	path[0] = rnd

	bigOne := big.NewInt(1)
	srvListLenWithoutLast := big.NewInt(int64(serversLen - 1))

	for i := 1; i < length; i++{

		rnd, err = rand.Int(rand.Reader, srvListLenWithoutLast)
		if err != nil {	return nil, err	}

		lastPlusOne := new(big.Int).Add(path[i-1], bigOne)
		rnd.Add(rnd, lastPlusOne)
		rnd.Mod(rnd, srvListLen)

		path[i] = rnd
	}

	pathStr := make([]string, len(path))

	for i := 0; i < len(path); i++{
		t, _ := strconv.Atoi(path[i].String())
		name := names[t]
		pathStr[i] = name
	}

	return pathStr, nil
}

//GetServerList returns a slice of all known server names
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

func (sli *ServerList)DebugGetServers(client bool){
	servPath := "servers/"
	ipPath := "/ip"
	var portPath string
	if client {	portPath = "/clientPort"
	}else {portPath = "/serverPort"}
	pubRSAPath := "/publicKeyRSA"
	pubElGamalPath := "/publicKeyElGamal"

	servers, err := ioutil.ReadDir(servPath)
	if err != nil {	log.Print(err); return }
	var name string
	var ip string
	var port string
	var pubRSA *rsa.PublicKey
	var pubElGamal *elgamal.PublicKey


	srvMap := make(map[string]*entry.Entry)


	for _, server := range servers {
		if server.IsDir(){
			name = server.Name()

			currPath := servPath + name

			ipFile, err := ioutil.ReadFile(currPath + ipPath)
			if err != nil {	log.Fatal(err) }

			ip = strings.TrimSpace(string(ipFile))

			portFile, err := ioutil.ReadFile(currPath + portPath)
			if err != nil {	log.Fatal(err) }

			port = strings.TrimSpace(string(portFile))

			pubRSA, err = crypt.LoadRSAPublic(currPath + pubRSAPath)
			if err != nil{
				log.Println(err)
				pubRSA = nil
			}

			privElGamal, err := crypt.LoadElGamal(currPath + pubElGamalPath)
			if err == nil{
				pubElGamal = &privElGamal.PublicKey
			}else {
				pubElGamal = nil
			}

			srvMap[name] = entry.New(name, ip + ":" + port, pubRSA, pubElGamal)
		}
	}

	sli.list = srvMap
}
package srvlist

import (
	"sync"
	"math/big"
	"crypto/rand"
	"errors"
	"crypt"
	"srvlist/entry"
	"strconv"
	"crypto/rsa"
	"golang.org/x/crypto/openpgp/elgamal"
	"io/ioutil"
	"log"
	"strings"
)

type ServerList struct{
	listmutex sync.Mutex
	list      map[string]*entry.Entry
}

func New()(*ServerList)  {
	sli := new(ServerList)

	sli.list = make(map[string]*entry.Entry)

	return sli
}

func (sli *ServerList)SetList(list map[string]*entry.Entry){
	sli.list = list;
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
	entr, ok := sli.list[name]
	sli.listmutex.Unlock()

	if ok{
		return entr.Encrypter, nil
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
	//no need for mutex, GetServerList is thread-safe
	names := sli.GetServerList()

	serversLen := len(names)

	if serversLen < 2{
		return nil, errors.New("too few servers to create a path");
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

func (sli *ServerList)DebugGetServers(){
	servPath := "servers/"
	ipportString := "/ipport"
	pubRSAString := "/publicKeyRSA"
	pubElGamalString := "/publicKeyElGamal"

	servers, err := ioutil.ReadDir(servPath)
	if err != nil {	log.Fatal(err) }
	var name string
	var ipport string
	var pubRSA *rsa.PublicKey
	var pubElGamal *elgamal.PublicKey

	currPath := servPath

	srvMap := make(map[string]*entry.Entry)


	for _, server := range servers {
		if server.IsDir(){
			name = server.Name()

			currPath += name

			ipportFile, err := ioutil.ReadFile(currPath + ipportString)
			if err != nil {	log.Fatal(err) }

			ipport = strings.TrimSpace(string(ipportFile))

			currPath = servPath + name

			pubRSA, err = crypt.LoadRSAPublic(currPath + pubRSAString)
			if err != nil{
				log.Println(err)
				pubRSA = nil
			}

			currPath = servPath + name

			privElGamal, err := crypt.LoadElGamal(currPath + pubElGamalString)
			if err == nil{
				pubElGamal = &privElGamal.PublicKey
			}else {
				pubElGamal = nil
			}

			srvMap[name] = entry.New(name, ipport, pubRSA, pubElGamal)

			currPath = servPath
		}
	}

	sli.list = srvMap
}
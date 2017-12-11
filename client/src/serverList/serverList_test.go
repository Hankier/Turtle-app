package serverList

import (
	"testing"
	"math/rand"
	"log"
)

func TestServerList_GetRandomPath(t *testing.T) {
	sli := NewServerList()
	var path []string
	sli.serverList = make(map[string]*serverEntry)

	path_len := 25;

	_, err := sli.GetRandomPath(path_len)

	if err != nil{
		log.Println(err)
	}else{
		t.Fail()
	}

	sli.serverList["0"] = NewServerEntry("0", "0", nil, nil)


	_, err = sli.GetRandomPath(path_len)

	if err != nil{
		log.Println(err)
	}else{
		t.Fail()
	}

	sli.serverList["1"] = NewServerEntry("1", "1", nil, nil)

	_, err = sli.GetRandomPath(path_len)

	if err != nil{
		log.Println(err)
	}else{
		t.Fail()
	}

	var kstring string

	for k := 2; k < 1000; k++{
		kstring = string(k)
		sli.serverList[kstring] = NewServerEntry(kstring, kstring, nil, nil)
		rnd := rand.Intn(path_len)
		path, err = sli.GetRandomPath(rnd)
		if rnd > 0{
			for i := 1; i < len(path); i++{
				if path[i] == path[i-1]{
					t.Fail()
				}
			}
		}
	}
}

func TestServerList_Getters(t *testing.T) {
}
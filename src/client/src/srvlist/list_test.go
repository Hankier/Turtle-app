package srvlist

import (
	"testing"
	"math/rand"
	"srvlist/entry"
	"strconv"
)

func TestServerList_GetRandomPath(t *testing.T) {
	sli := New()
	var path []string
	sli.list = make(map[string]*entry.Entry)

	pathLen := 25;

	_, err := sli.GetRandomPath(pathLen)

	if err == nil{
		t.Fail()
	}

	sli.list["0"] = entry.New("0", "0", nil, nil)


	_, err = sli.GetRandomPath(pathLen)

	if err == nil{
		t.Fail()
	}

	sli.list["1"] = entry.New("1", "1", nil, nil)

	_, err = sli.GetRandomPath(pathLen)

	if err != nil{
		t.Fail()
	}

	var kstring string

	for k := 2; k < 1000; k++{
		kstring = strconv.Itoa(k)
		sli.list[kstring] = entry.New(kstring, kstring, nil, nil)
		rnd := rand.Intn(pathLen)
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
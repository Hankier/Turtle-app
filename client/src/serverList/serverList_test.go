package serverList

import (
	"testing"
	"fmt"
)

func TestServerList_GetRandomPath(t *testing.T) {
	sli := NewServerList()
	path1 := sli.GetRandomPath(5)
	path2 := sli.GetRandomPath(5)

	count := 0

	for i := 0; i < 5; i++{
		fmt.Println(path1[i] + "::" + path2[i])
		if path1[i] == path2[i]{
			count++
		}
	}

	if count == 5{
		t.Error("Got two identical paths")
	}
}
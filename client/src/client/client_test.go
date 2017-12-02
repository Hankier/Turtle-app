package client

import (
	"testing"
	"fmt"
)

func TestClient_GetRandomPath(t *testing.T) {
	client := NewClient()
	path1 := client.GetRandomPath(5)
	path2 := client.GetRandomPath(5)

	count := 0

	for i := 0; i < 5; i++{
		fmt.Println(path1[i].Name + "::" + path2[i].Name)
		if path1[i].Name == path2[i].Name{
			count++
		}
	}

	if count == 5{
		t.Error("Got two identical paths")
	}
}
package serverlist

import "testing"

func TestServerList_AddServerToList(t *testing.T){
	srv := NewList()
	new_ip := "55.55.55.55"
	new_key := "5a53a313ad13f"
	srv.AddServerToList(new_ip, new_key)

	added_ip, _ := srv.GetServerIpPort(srv.GetServerList()[0])
	if added_ip != new_ip {
		t.Error("Expected ",new_ip, " got ", added_ip)
	}

}

func TestServerList_RemoveServerFromList(t *testing.T){
	srv := NewList()
	new_ip := "55.55.55.55"
	new_key := "5a53a313ad13f"
	srv.AddServerToList(new_ip, new_key)

	srv.RemoveServerFromList("00000000")

	list_size := len(srv.GetServerList())

	if list_size != 0 {
		t.Error("Expected 0 got ", list_size)
	}

}

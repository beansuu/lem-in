package ant

import "testing"

func TestAddRoom(t *testing.T) {
	antNest := AntNest{0, make(map[string][2]int), make(map[string][]string), "", ""}
	antNest.addRoom("room1", 1, 2)
	if antNest.rooms["room1"] != [2]int{1, 2} {
		t.Error("addRoom failed")
	}
}

func TestAddTunnel(t *testing.T) {
	antNest := AntNest{0, make(map[string][2]int), make(map[string][]string), "", ""}
	antNest.addTunnel("room1", "room2")
	if len(antNest.tunnels["room1"]) != 1 || antNest.tunnels["room1"][0] != "room2" {
		t.Error("addTunnel failed")
	}
	if len(antNest.tunnels["room2"]) != 1 || antNest.tunnels["room2"][0] != "room1" {
		t.Error("addTunnel failed")
	}
}

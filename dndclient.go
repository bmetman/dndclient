package dndclient

import "dndclient/data/gateway"

func GetMonsters() string {
	return gateway.Get("monsters")
}

package main

import (
	"dndclient/data"
	"fmt"
)

func main() {
	fmt.Println(data.Get("monsters"))
}

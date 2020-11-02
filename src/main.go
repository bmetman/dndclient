package main

import (
	"fmt"
)

func sum(x int, y int) int {
	return x + y
}

func main() {
	fmt.Printf("%d", sum(5, 5))
}

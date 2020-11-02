package main

import (
	"fmt"
)

func Sum(x int, y int) int {
	return x + y
}

func main() {
	fmt.Printf("%d", Sum(5, 5))
}

package main

import (
	"fmt"
)

func main() {
	var N int

	fmt.Scan(&N)

	fmt.Printf("%d", 1<<N)
}

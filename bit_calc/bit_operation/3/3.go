package main

import (
	"fmt"
)

func main() {
	var N int
	var M int

	fmt.Scan(&N)
	fmt.Scan(&M)

	fmt.Printf("%d", N>>M)
}

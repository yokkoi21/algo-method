package main

import (
	"fmt"
)

func main() {
	var A int
	var i int

	fmt.Scan(&A)
	fmt.Scan(&i)

	fmt.Printf("%d\n", A|(1<<i))
}

package main

import (
	"fmt"
)

func main() {
	var N int
	var A int

	fmt.Scan(&N)
	fmt.Scan(&A)

	if N%A == 0 {
		fmt.Println(N / A)
	} else {
		fmt.Println(N/A + 1)
	}
}

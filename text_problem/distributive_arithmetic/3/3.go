package main

import (
	"fmt"
)

func main() {
	var A int
	var K int

	fmt.Scan(&A)
	fmt.Scan(&K)

	if A%(K+1) != 0 {
		fmt.Println(-1)
	} else {
		fmt.Println((A / (K + 1)) * K)
	}
}

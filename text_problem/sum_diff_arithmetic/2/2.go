package main

import (
	"fmt"
)

func main() {
	var A int

	fmt.Scan(&A)

	if A%2 != 0 {
		fmt.Println(-1)
	} else {
		fmt.Println(A / 2)
	}
}

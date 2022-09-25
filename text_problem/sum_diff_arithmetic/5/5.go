package main

import (
	"fmt"
)

func main() {
	var A int
	var B int

	fmt.Scan(&A)
	fmt.Scan(&B)

	if (A+B)%2 != 0 || A < B {
		fmt.Println(-1)
	} else {
		fmt.Println((A + B) / 2)
	}
}

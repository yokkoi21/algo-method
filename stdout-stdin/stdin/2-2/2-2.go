package main

import (
	"fmt"
)

func main() {
	var A int
	var B int

	fmt.Scan(&A)
	fmt.Scan(&B)
	if A > B {
		fmt.Println(A)
	} else {
		fmt.Println(B)
	}
}

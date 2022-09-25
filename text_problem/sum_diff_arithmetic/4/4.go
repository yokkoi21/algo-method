package main

import (
	"fmt"
)

func main() {
	var A int
	var B int

	fmt.Scan(&A)
	fmt.Scan(&B)

	if (A+B)%2 == 0 {
		fmt.Println((A + B) / 2)
	} else {
		fmt.Println(-1)
	}
}

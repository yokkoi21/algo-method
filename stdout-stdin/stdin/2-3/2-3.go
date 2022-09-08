package main

import (
	"fmt"
)

func main() {
	var A int
	var B int

	fmt.Scan(&A)
	fmt.Scan(&B)
	if A%10 > B%10 {
		fmt.Println(B)
	} else {
		fmt.Println(A)
	}
}

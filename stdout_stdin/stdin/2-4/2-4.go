package main

import (
	"fmt"
)

func main() {
	var A int
	var B int

	fmt.Scan(&A)
	fmt.Scan(&B)
	if A%B == 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

package main

import (
	"fmt"
)

func main() {
	var A int
	var B int

	fmt.Scan(&A)
	fmt.Scan(&B)

	fmt.Printf("%d\n", A&B)
	fmt.Printf("%d\n", A|B)
}

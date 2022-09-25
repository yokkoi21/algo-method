package main

import (
	"fmt"
)

func main() {
	var A int
	var B int
	var K int

	fmt.Scan(&A)
	fmt.Scan(&B)
	fmt.Scan(&K)

	fmt.Println((B / K) - ((A - 1) / K))

}

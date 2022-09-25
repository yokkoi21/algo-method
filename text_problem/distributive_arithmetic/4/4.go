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

	if (B-A)%(K+1) != 0 || B < A {
		fmt.Println(-1)
	} else {
		fmt.Println((((B - A) / (K + 1)) * K) + A)
	}
}

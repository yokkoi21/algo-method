package main

import (
	"fmt"
)

func main() {
	var A int
	var B int
	var K int
	var x int

	fmt.Scan(&A)
	fmt.Scan(&B)
	fmt.Scan(&K)

	if (B+A)%(K+1) != 0 {
		fmt.Println(-1)
	} else {
		x = (((B + A) / (K + 1)) * K) - A
		if x < 0 {
			fmt.Println(-1)
		} else {
			fmt.Println(x)
		}
	}
}

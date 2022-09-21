package main

import (
	"fmt"
)

func main() {
	var A int
	var K int

	fmt.Scan(&A)
	fmt.Scan(&K)

	for i := A; i < A+K; i++ {
		if i%K == 0 {
			fmt.Println(i)
			break
		}
	}
}

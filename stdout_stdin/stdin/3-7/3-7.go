package main

import (
	"fmt"
)

func main() {
	var N int
	var A []int
	var tmp int
	var min int = 1000

	fmt.Scan(&N)
	for i := 0; i < N; i++ {
		fmt.Scan(&tmp)
		A = append(A, tmp)
	}

	for i := 0; i < N; i++ {
		if min > A[i] {
			min = A[i]
		}
	}
	fmt.Println(min)
}

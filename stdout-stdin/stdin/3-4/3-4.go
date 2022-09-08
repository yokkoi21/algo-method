package main

import (
	"fmt"
)

func main() {
	var N int
	var A []int
	var tmp int

	fmt.Scan(&N)
	for i := 0; i < N; i++ {
		fmt.Scan(&tmp)
		A = append(A, tmp)
	}

	for i := 0; i < N; i++ {
		if A[i]%3 == 0 {
			fmt.Println(A[i])
		}
	}
}

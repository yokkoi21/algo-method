package main

import (
	"fmt"
)

func main() {
	var N int
	var A []string
	var tmp string
	var ans string

	fmt.Scan(&N)
	for i := 0; i < N; i++ {
		fmt.Scan(&tmp)
		A = append(A, tmp)
	}

	for i := 0; i < N; i++ {
		ans += A[i]
	}
	fmt.Println(len(ans))
}

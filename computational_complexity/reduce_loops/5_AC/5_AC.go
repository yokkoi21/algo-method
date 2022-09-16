package main

import (
	"fmt"
)

func main() {
	var N int
	var tmp int
	var A []int
	var sum int
	var squared_sum int

	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		fmt.Scan(&tmp)
		A = append(A, tmp)
		sum += tmp
	}

	for i := 0; i < N; i++ {
		squared_sum += A[i] * A[i]
	}

	fmt.Println(((sum * sum) - squared_sum) / 2)
}

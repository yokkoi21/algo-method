package main

import (
	"fmt"
)

func main() {
	var N int
	//var V int
	var tmp int
	var A []int
	var max int
	var max_i int

	fmt.Scan(&N)
	//fmt.Scan(&V)
	for i := 0; i < N; i++ {
		fmt.Scan(&tmp)
		A = append(A, tmp)
	}

	for i := 0; i < N; i++ {
		if A[i] > max {
			//fmt.Println(A[i])
			max = A[i]
			max_i = i
		}
	}a
	fmt.Println(max_i)

}

package main

import (
	"fmt"
)

func main() {
	var N int
	//var V int
	var tmp int
	var A []int
	var min int = 100
	//var max_i int

	fmt.Scan(&N)
	//fmt.Scan(&V)
	for i := 0; i < N; i++ {
		fmt.Scan(&tmp)
		A = append(A, tmp)
	}

	for i := 0; i < N; i++ {
		if A[i] < min {
			//fmt.Println(A[i])
			min = A[i]
			//max_i = i
		}
	}
	fmt.Println(min)

}

package main

import (
	"fmt"
)

func main() {
	var N int
	//var V int
	var tmp int
	var A []int
	var cnt int

	fmt.Scan(&N)
	//fmt.Scan(&V)
	for i := 0; i < N; i++ {
		fmt.Scan(&tmp)
		A = append(A, tmp)
	}

	for i := 1; i < N; i++ {
		if A[i] > A[i-1] {
			//fmt.Println(A[i])
			cnt++
		}
	}
	fmt.Println(cnt)

}

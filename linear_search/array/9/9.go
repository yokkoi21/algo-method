package main

import (
	"fmt"
)

func main() {
	var N int
	//var V int
	var tmp int
	var A []int
	var num_cnt [10]int
	//var max_i int

	fmt.Scan(&N)
	//fmt.Scan(&V)
	for i := 0; i < N; i++ {
		fmt.Scan(&tmp)
		A = append(A, tmp)
	}

	for i := 0; i < N; i++ {
		num_cnt[A[i]-1]++
	}
	for i := 0; i < 9; i++ {
		fmt.Println(num_cnt[i])
	}

}

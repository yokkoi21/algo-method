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
	var num_cnt_max int
	var num_cnt_max_i int
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
		if num_cnt[i] > num_cnt_max {
			num_cnt_max = num_cnt[i]
			num_cnt_max_i = i + 1
		}
	}

	fmt.Println(num_cnt_max_i)
}

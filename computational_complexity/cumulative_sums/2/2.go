package main

import (
	"fmt"
)

func main() {
	var N int
	var Q int
	var tmp_l int
	var tmp_r int
	var sum_slice []int
	var tmp int
	var A []int

	fmt.Scan(&N)

	sum_slice = append(sum_slice, 0)
	for i := 0; i < N; i++ {
		fmt.Scan(&tmp)
		A = append(A, tmp)
		sum_slice = append(sum_slice, sum_slice[i]+A[i])
	}

	fmt.Scan(&Q)
	//fmt.Println(A)
	//fmt.Println(sum_slice)

	for i := 0; i < Q; i++ {
		fmt.Scan(&tmp_l)
		fmt.Scan(&tmp_r)
		fmt.Println(sum_slice[tmp_r] - sum_slice[tmp_l])
	}

}

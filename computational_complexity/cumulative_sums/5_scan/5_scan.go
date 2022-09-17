package main

import (
	"fmt"
)

func main() {
	var N int
	var Q int
	var tmp_l int
	var tmp_r int
	var num_slice [100002]int
	var sum_slice [100002]int
	var tmp int
	var L []int

	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		fmt.Scan(&tmp)
		L = append(L, tmp)
		num_slice[L[i]]++
	}

	for i := 0; i < 100001; i++ {
		sum_slice[i+1] = sum_slice[i] + num_slice[i]
	}

	fmt.Scan(&Q)

	for i := 0; i < Q; i++ {
		fmt.Scan(&tmp_l)
		fmt.Scan(&tmp_r)
		fmt.Println(sum_slice[tmp_r+1] - sum_slice[tmp_l])
	}

}

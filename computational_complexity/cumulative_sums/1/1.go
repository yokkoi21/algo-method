package main

import (
	"fmt"
)

func main() {
	var N int
	var Q int
	var tmp_k int
	var sum_slice []int
	var tmp int
	var sum int
	//var A []int

	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		fmt.Scan(&tmp)
		//A = append(A, tmp)
		sum += tmp
		sum_slice = append(sum_slice, sum)
	}

	fmt.Scan(&Q)

	for i := 0; i < Q; i++ {
		fmt.Scan(&tmp_k)
		if tmp_k != 0 {
			fmt.Println(sum_slice[tmp_k-1])
		} else {
			fmt.Println(0)
		}

	}

}

package main

import (
	"fmt"
)

func main() {
	var N int
	var tmp int
	var A []int
	var is_prime bool = true
	var cnt int

	fmt.Scan(&N)
	for i := 0; i < N; i++ {
		fmt.Scan(&tmp)
		A = append(A, tmp)
	}
	//fmt.Println(len(A))

	//fmt.Println(" ")
	for i := 0; i < N; i++ {
		if A[i] == 1 {
			is_prime = false
		}
		for j := 2; j < A[i]; j++ {
			if A[i]%j == 0 {
				is_prime = false
			}
		}
		if is_prime {

			cnt++
		}
		is_prime = true
	}

	fmt.Println(cnt)
}

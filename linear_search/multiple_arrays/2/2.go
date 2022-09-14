package main

import (
	"fmt"
)

func main() {
	var N int
	var M int
	var K int
	var tmp int
	var A []int
	var B []int
	var cnt int

	fmt.Scan(&N)
	fmt.Scan(&M)
	fmt.Scan(&K)

	for i := 0; i < N; i++ {
		fmt.Scan(&tmp)
		A = append(A, tmp)
	}
	for i := 0; i < M; i++ {
		fmt.Scan(&tmp)
		B = append(B, tmp)
	}

	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			if A[i]+B[j] == K {
				cnt++
			}
		}
	}

	fmt.Println(cnt)
}

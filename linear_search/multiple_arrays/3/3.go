package main

import (
	"fmt"
)

func main() {
	var X int
	var Y int
	var Z int
	var tmp int
	var A []int
	var B []int
	var C []int
	var cnt int

	fmt.Scan(&X)
	fmt.Scan(&Y)
	fmt.Scan(&Z)

	for i := 0; i < X; i++ {
		fmt.Scan(&tmp)
		A = append(A, tmp)
	}
	for i := 0; i < Y; i++ {
		fmt.Scan(&tmp)
		B = append(B, tmp)
	}
	for i := 0; i < Z; i++ {
		fmt.Scan(&tmp)
		C = append(C, tmp)
	}

	for i := 0; i < X; i++ {
		for j := 0; j < Y; j++ {
			for k := 0; k < Z; k++ {
				if A[i]+B[j] == C[k] {
					cnt++
				}
			}
		}
	}

	fmt.Println(cnt)
}

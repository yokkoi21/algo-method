package main

import (
	"fmt"
)

func main() {
	var N int
	var K int
	var tmp int
	var A []int
	var cnt int

	fmt.Scan(&N)
	fmt.Scan(&K)

	for i := 0; i < N; i++ {
		fmt.Scan(&tmp)
		A = append(A, tmp)
	}

	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			// fmt.Print("i:")
			// fmt.Println(i)
			// fmt.Print("j:")
			// fmt.Println(j)
			if A[i]+A[j] <= K {
				cnt++
			}
		}
	}

	fmt.Println(cnt)
}

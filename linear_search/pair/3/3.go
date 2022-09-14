package main

import (
	"fmt"
)

func main() {
	var N int
	var tmp int
	var A []int
	var cnt int

	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		fmt.Scan(&tmp)
		A = append(A, tmp)
	}

	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			for k := j + 1; k < N; k++ {
				// fmt.Print("i:")
				// fmt.Print(i)
				// fmt.Print("j:")
				// fmt.Print(j)
				// fmt.Print("k:")
				// fmt.Println(k)
				if A[i] <= A[j] && A[j] >= A[k] {
					cnt++
				}
			}
		}
	}

	fmt.Println(cnt)
}

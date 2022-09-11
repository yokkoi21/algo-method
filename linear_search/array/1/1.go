package main

import (
	"fmt"
)

func main() {
	var N int
	var V int
	var tmp int
	var A []int
	var flag bool = false

	fmt.Scan(&N)
	fmt.Scan(&V)
	for i := 0; i < N; i++ {
		fmt.Scan(&tmp)
		A = append(A, tmp)
	}

	for i := 0; i < N; i++ {
		if A[i] == V {
			flag = true
		}
	}
	if flag {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}

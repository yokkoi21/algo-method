package main

import (
	"fmt"
)

func main() {
	var N int
	var A []string
	var tmp string
	var l_cnt int
	var r_cnt int

	fmt.Scan(&N)
	for i := 0; i < N; i++ {
		fmt.Scan(&tmp)
		A = append(A, tmp)
	}

	for i := 0; i < N; i++ {
		if A[i] == "right" {
			r_cnt += 1
		} else if A[i] == "left" {
			l_cnt += 1
		}
	}
	if r_cnt == l_cnt {
		fmt.Println("same")
	} else if r_cnt > l_cnt {
		fmt.Println("right")
	} else if r_cnt < l_cnt {
		fmt.Println("left")
	}
}

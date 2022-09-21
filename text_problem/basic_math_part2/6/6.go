package main

import (
	"fmt"
)

func main() {
	var U int
	var T int
	var A int
	var sum_T int

	fmt.Scan(&U)
	fmt.Scan(&T)
	fmt.Scan(&A)

	if (A-U)%T == 0 {
		fmt.Println(A)
	} else {
		for sum_T < (A - U) {
			sum_T += T
		}
		fmt.Println(U + sum_T)
	}
}

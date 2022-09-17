package main

import (
	"fmt"
)

func main() {
	var N int
	var D int
	var sum_slice []int
	var tmp int
	var X []int
	var max int
	var ans_i int

	fmt.Scan(&N)
	fmt.Scan(&D)

	sum_slice = append(sum_slice, 0)
	for i := 0; i < N; i++ {
		fmt.Scan(&tmp)
		X = append(X, tmp)
		sum_slice = append(sum_slice, sum_slice[i]+X[i])
	}
	//fmt.Print(sum_slice)
	//fmt.Print(X)

	for i := 0; i < N-D+1; i++ {
		if sum_slice[i+D]-sum_slice[i] >= max {
			max = sum_slice[i+D] - sum_slice[i]
			ans_i = i
		}
	}

	fmt.Print(ans_i)
	fmt.Print("~")
	fmt.Println(ans_i + D - 1)

}

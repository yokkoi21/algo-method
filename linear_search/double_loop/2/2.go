package main

import (
	"fmt"
)

func main() {
	var N int
	var K int
	var factor_cnt int
	var ans_cnt int

	fmt.Scan(&N)
	fmt.Scan(&K)

	for i := 1; i < N+1; i++ {
		for j := 1; j < i+1; j++ {
			if i%j == 0 {
				factor_cnt++
			}
		}
		//fmt.Println(factor_cnt)
		if factor_cnt == K {
			ans_cnt++
		}
		factor_cnt = 0
	}

	fmt.Println(ans_cnt)
}

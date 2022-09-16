package main

import (
	"fmt"
)

func main() {
	var N int
	var M int
	var cnt int

	fmt.Scan(&N)
	fmt.Scan(&M)

	for i := 1; i < N+1; i++ {
		for j := 1; j < N+1; j++ {
			for k := 1; k < N+1; k++ {
				if i+j+k <= M {
					cnt++
				}
			}
		}
	}

	fmt.Println(cnt)
}

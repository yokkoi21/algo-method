package main

import (
	"fmt"
)

func main() {
	var N int
	var cnt int

	fmt.Scan(&N)

	for i := 1; i < N+1; i++ {
		if i%2 != 0 && i%3 != 0 && i%5 != 0 {
			cnt++
		}
	}

	fmt.Println(cnt)
}

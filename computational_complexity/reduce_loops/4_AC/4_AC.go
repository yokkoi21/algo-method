package main

import (
	"fmt"
)

func main() {
	var N int
	var tmp int
	var sum int

	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		fmt.Scan(&tmp)
		sum += tmp
	}

	fmt.Println(sum * sum)
}

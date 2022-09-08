package main

import (
	"fmt"
)

func main() {
	var tmp int
	var num []int
	var max int

	for i := 0; i < 4; i++ {
		fmt.Scan(&tmp)
		num = append(num, tmp)
	}

	for i := 0; i < 4; i++ {
		if max < num[i] {
			max = num[i]
		}
	}
	fmt.Println(max)
}

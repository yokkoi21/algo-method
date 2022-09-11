package main

import (
	"fmt"
)

func main() {
	var A int
	var B int
	var tmp int
	var gcd int

	fmt.Scan(&A)
	fmt.Scan(&B)

	if A > B {
		tmp = A
	} else {
		tmp = B
	}

	for i := 1; i < tmp+1; i++ {
		if A%i == 0 && B%i == 0 {
			gcd = i
		}
	}

	fmt.Println(gcd)

}

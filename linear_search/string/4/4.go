package main

import (
	"fmt"
)

func main() {
	var N int
	var S string
	var T string
	var cnt int

	fmt.Scan(&N)
	fmt.Scan(&S)
	fmt.Scan(&T)

	for i := 0; i < len(S); i++ {
		if S[i] != T[i] {
			cnt++
		}
	}

	fmt.Println(cnt)

}

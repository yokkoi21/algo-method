package main

import (
	"fmt"
)

func main() {
	var S string
	var N int
	var M int

	fmt.Scan(&S)
	fmt.Scan(&N)
	fmt.Scan(&M)

	for i := 0; i < len(S); i++ {
		if i == N-1 {
			fmt.Print(string(S[M-1]))
		} else if i == M-1 {
			fmt.Print(string(S[N-1]))
		} else {
			fmt.Print(string(S[i]))
		}
	}
}

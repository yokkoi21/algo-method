package main

import (
	"fmt"
)

func main() {
	var N int
	var S string
	var cnt int

	fmt.Scan(&N)
	fmt.Scan(&S)

	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			fmt.Print("i:")
			fmt.Print(i)
			fmt.Print("j:")
			fmt.Println(j)
			if S[i] == S[j] {
				cnt++
			}
		}
	}

	fmt.Println(cnt)
}

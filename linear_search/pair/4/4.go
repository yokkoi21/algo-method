package main

import (
	"fmt"
)

func main() {
	var N int
	var tmp string
	var S []string

	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		fmt.Scan(&tmp)
		S = append(S, tmp)
	}

	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			// fmt.Print("i:")
			// fmt.Print(i)
			// fmt.Print("j:")
			// fmt.Println(j)
			if S[i] == S[j] {
				fmt.Println("Yes")
				return
			}
		}
	}

	fmt.Println("No")
}

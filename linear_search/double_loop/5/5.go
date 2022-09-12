package main

import (
	"fmt"
)

func main() {
	var N int
	var tmp string
	var S []string
	var palindrome_flag bool = true
	var cnt int

	fmt.Scan(&N)
	for i := 0; i < N; i++ {
		fmt.Scan(&tmp)
		S = append(S, tmp)
	}

	for i := 0; i < N; i++ {
		for j := 0; j < len(S[i]); j++ {
			if S[i][j] != S[i][len(S[i])-j-1] {
				palindrome_flag = false
			}
		}
		if palindrome_flag {
			cnt++
		}
		palindrome_flag = true
	}

	fmt.Println(cnt)
}

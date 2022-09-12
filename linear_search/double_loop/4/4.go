package main

import (
	"fmt"
)

func main() {
	var S string
	var alphabet [26]int
	var cnt int

	fmt.Scan(&S)

	for i := 0; i < len(S); i++ {
		for j := 0; j < 26; j++ {
			if int(S[i]) == (97 + j) {
				alphabet[j] = 1
			}
		}
	}

	for i := 0; i < 26; i++ {
		if alphabet[i] == 1 {
			cnt++
		}
	}
	fmt.Println(cnt)
}

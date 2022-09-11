package main

import (
	"fmt"
)

func main() {
	var S string
	var cnt int

	fmt.Scan(&S)

	for i := 1; i < len(S); i++ {
		if S[i] == S[i-1] {
			cnt++
		}
	}

	fmt.Println(cnt)

}

package main

import (
	"fmt"
)

func main() {
	var L int
	var R int
	var cnt int

	fmt.Scan(&L)
	fmt.Scan(&R)

	for i := L; i < R+1; i++ {
		for j := i + 1; j < R+1; j++ {
			// fmt.Print("i:")
			// fmt.Print(i)
			// fmt.Print("j:")
			// fmt.Println(j)
			if i%10 == j%10 {
				cnt++
			}
		}
	}

	fmt.Println(cnt)
}

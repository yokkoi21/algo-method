package main

import (
	"fmt"
)

func main() {
	var S string
	var c string
	var flag bool

	fmt.Scan(&S)
	fmt.Scan(&c)

	for i := 0; i < len(S); i++ {
		if string(S[i]) == c {
			flag = true
		}
	}
	if flag {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}

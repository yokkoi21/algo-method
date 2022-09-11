package main

import (
	"fmt"
)

func main() {
	var S string
	var flag bool = true

	fmt.Scan(&S)

	for i := 0; i < len(S); i++ {
		if string(S[i]) != string(S[len(S)-1-i]) {
			flag = false
		}
	}
	if flag {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}

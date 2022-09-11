package main

import (
	"fmt"
)

func main() {
	var S string
	var T string

	fmt.Scan(&S)
	fmt.Scan(&T)
	if S == T {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

package main

import (
	"fmt"
	"strings"
)

func main() {
	var S string
	var T string

	fmt.Scan(&S)
	fmt.Scan(&T)

	if strings.Contains(S, T) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}

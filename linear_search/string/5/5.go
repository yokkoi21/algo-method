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

	if strings.Index(S, T) != -1 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}

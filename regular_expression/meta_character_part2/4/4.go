package main

import (
	"fmt"
	"regexp"
)

func main() {
	var S string

	fmt.Scan(&S)

	r := regexp.MustCompile(`^[a-zA-Z0-9]+@[a-zA-Z0-9]+\.[a-zA-Z0-9]{1,4}$`)

	if r.MatchString(S) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

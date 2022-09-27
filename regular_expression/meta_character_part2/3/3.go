package main

import (
	"fmt"
	"regexp"
)

func main() {
	var S string

	fmt.Scan(&S)

	r := regexp.MustCompile(`^\d{3}-\d{4}$`)

	if r.MatchString(S) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

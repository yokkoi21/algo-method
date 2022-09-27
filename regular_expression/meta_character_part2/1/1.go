package main

import (
	"fmt"
	"regexp"
)

func main() {
	var S string

	fmt.Scan(&S)

	r := regexp.MustCompile(`\d`)

	if r.MatchString(S) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

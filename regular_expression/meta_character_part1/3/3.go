package main

import (
	"fmt"
	"regexp"
)

func main() {
	var S string
	//var R string

	fmt.Scan(&S)

	r := regexp.MustCompile(`^a{1,5}b{10}c*$`)

	if r.MatchString(S) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

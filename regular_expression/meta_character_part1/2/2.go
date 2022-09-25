package main

import (
	"fmt"
	"regexp"
)

func main() {
	var S string
	//var R string

	fmt.Scan(&S)

	r := regexp.MustCompile(`^metho+d$`)

	if r.MatchString(S) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

package main

import (
	"fmt"
	"regexp"
)

func main() {
	var S string

	fmt.Scan(&S)

	r := regexp.MustCompile("rur")

	for r.MatchString(S) {
		S = r.ReplaceAllString(S, "rar")
	}
	fmt.Println(S)
}

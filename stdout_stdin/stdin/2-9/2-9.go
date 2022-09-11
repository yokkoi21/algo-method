package main

import (
	"fmt"
)

func main() {
	var S string
	var N int

	fmt.Scan(&S)
	fmt.Scan(&N)

	fmt.Println(string(S[N-1]))

}

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var N int
	var S string
	var B int64

	fmt.Scan(&N)

	S += "1"

	for i := 0; i < N; i++ {
		S = S + "0"
	}
	B, _ = strconv.ParseInt(S, 2, 0)
	fmt.Println(B)
}

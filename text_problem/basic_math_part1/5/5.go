package main

import (
	"fmt"
)

func main() {
	var A int
	var B int
	var C int
	var D int

	fmt.Scan(&A)
	fmt.Scan(&B)
	fmt.Scan(&C)
	fmt.Scan(&D)

	fmt.Println(A*C + B*D)
}

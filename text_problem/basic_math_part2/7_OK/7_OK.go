package main

import (
	"fmt"
)

func main() {
	var A int
	var K int

	fmt.Scan(&A)
	fmt.Scan(&K)

	fmt.Println((A + K - 1) / K * K)

}

package main

import (
	"fmt"
)

func main() {
	var M int
	var K int

	fmt.Scan(&M)
	fmt.Scan(&K)

	if M%K == 0 {
		fmt.Println(M / K)
	} else {
		fmt.Println(M/K + 1)
	}
}

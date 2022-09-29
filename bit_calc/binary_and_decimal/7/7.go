package main

import (
	"fmt"
	"math"
)

func main() {
	var N float64
	//var B int64

	fmt.Scan(&N)

	fmt.Printf("%b", int(math.Pow(2, N)))
}

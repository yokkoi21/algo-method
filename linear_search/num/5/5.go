package main

import (
	"fmt"
)

func main() {
	var N int
	//var cnt int

	fmt.Scan(&N)

	for i := 1; i < N+1; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}

}

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var A int
	//var i float64
	var i int

	fmt.Fscan(in, &A)
	fmt.Fscan(in, &i)

	//fmt.Println(A ^ int(math.Pow(2, i)))
	fmt.Println(A ^ (1 << i))
}

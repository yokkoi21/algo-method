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

	var X [8]int

	for i := 0; i < 8; i++ {
		fmt.Fscan(in, &X[i])
	}

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {

		}
	}

	fmt.Println(tmp1 ^ tmp2)
}

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
		for j := 15; j >= 0; j -= 2 {
			if (X[i]&(1<<j)) == 0 && (X[i]&(1<<(j-1))) == 0 {
				fmt.Print(".")
			} else if (X[i]&(1<<j)) == 0 && (X[i]&(1<<(j-1))) > 0 {
				fmt.Print("o")
			} else if (X[i]&(1<<j)) > 0 && (X[i]&(1<<(j-1))) == 0 {
				fmt.Print("x")
			}
		}
		fmt.Print("\n")
	}

}

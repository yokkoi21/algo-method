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

	var N int
	var S int
	var Q int
	var query int
	var x int

	fmt.Fscan(in, &N)
	fmt.Fscan(in, &S)
	fmt.Fscan(in, &Q)

	for i := 0; i < Q; i++ {
		fmt.Fscan(in, &query)
		fmt.Fscan(in, &x)
		if query == 0 {
			S = S ^ (1 << x)
		} else if query == 1 {
			if S&(1<<x) > 0 {
				fmt.Println("on")
			} else {
				fmt.Println("off")
			}
		}

	}

}

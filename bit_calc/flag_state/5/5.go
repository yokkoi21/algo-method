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
	var X int

	fmt.Fscan(in, &N)
	fmt.Fscan(in, &X)
	if (N & (1 << X)) > 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}

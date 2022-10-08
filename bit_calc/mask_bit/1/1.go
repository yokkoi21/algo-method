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
	var M int

	fmt.Fscan(in, &A)
	fmt.Fscan(in, &M)

	fmt.Fprintln(out, A|M)

}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var N int
	var cnt int
	var ans string

	fmt.Fscan(in, &N)

	for i := 0; i < 30; i++ {
		if (N & (1 << i)) > 0 {
			cnt++
			tmp := strconv.Itoa(i)
			ans += tmp + " "
		}
	}
	fmt.Fprintln(out, cnt)
	fmt.Fprintln(out, ans)

}

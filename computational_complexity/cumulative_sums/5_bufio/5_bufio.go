package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush() // バッファリングしているので、最後に確実に Flush して出力させます。

	var N int
	var Q int
	var tmp_l int
	var tmp_r int
	var num_slice [100002]int
	var sum_slice [100002]int
	var tmp int
	var L []int

	fmt.Fscan(in, &N)

	for i := 0; i < N; i++ {
		fmt.Fscan(in, &tmp)
		L = append(L, tmp)
		num_slice[L[i]]++
	}

	for i := 0; i < 100001; i++ {
		sum_slice[i+1] = sum_slice[i] + num_slice[i]
	}

	fmt.Fscan(in, &Q)

	for i := 0; i < Q; i++ {
		fmt.Fscan(in, &tmp_l)
		fmt.Fscan(in, &tmp_r)
		fmt.Fprintln(out, sum_slice[tmp_r+1]-sum_slice[tmp_l])
	}

}

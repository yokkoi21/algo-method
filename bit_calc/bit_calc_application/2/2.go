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
	var M int
	var led_statuses [10]string = [10]string{"1110111", "0100100", "1011101", "1101101", "0101110", "1101011", "1111011", "0100111", "1111111", "1101111"}

	fmt.Fscan(in, &N)
	fmt.Fscan(in, &M)

	tmp1, _ := strconv.ParseInt(led_statuses[N], 2, 0)
	tmp2, _ := strconv.ParseInt(led_statuses[M], 2, 0)

	fmt.Println(tmp1 ^ tmp2)
}

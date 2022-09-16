package main

import (
	"fmt"
	"math"
)

func main() {
	var N int
	var M int
	var tmp_z int
	var cnt int

	fmt.Scan(&N)
	fmt.Scan(&M)

	for i := 1; i < N+1; i++ {
		for j := 1; j < N+1; j++ {
			tmp_z = int(math.Min(float64(N), float64(M-i-j)))
			if tmp_z < 0 {
				continue
			}
			cnt += tmp_z
		}
	}

	fmt.Println(cnt)
}

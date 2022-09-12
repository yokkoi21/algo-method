package main

import (
	"fmt"
	"strconv"
)

func main() {
	var L int
	var R int
	var tmp string
	var flag bool = true
	var cnt int

	fmt.Scan(&L)
	fmt.Scan(&R)

	for i := L; i < R+1; i++ {
		tmp = strconv.Itoa(i)
		for j := 0; j < len(tmp); j++ {
			if tmp[j] != tmp[len(tmp)-j-1] {
				flag = false
			}
		}
		if flag {
			cnt++
		}
		flag = true
	}

	fmt.Println(cnt)
}

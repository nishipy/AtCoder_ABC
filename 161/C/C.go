package main

import (
	"fmt"
)

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	r := n % k

	fmt.Println(min(r, k-r))
}

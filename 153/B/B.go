package main

import (
	"fmt"
)

func main() {
	var h, n int
	fmt.Scan(&h, &n)
	sum := 0
	for i := 0; i < n; i++ {
		var tmp int
		fmt.Scan(&tmp)
		sum += tmp
	}
	if sum >= h {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

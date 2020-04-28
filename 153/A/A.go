package main

import (
	"fmt"
)

func main() {
	var h, a int
	fmt.Scan(&h, &a)

	ans := h / a
	if h%a != 0 {
		ans++
	}
	fmt.Println(ans)
}

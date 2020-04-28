package main

import (
	"fmt"
)

func main() {
	var n, a, b int
	fmt.Scan(&n, &a, &b)

	if n%(a+b) > a {
		fmt.Println(a*(n/(a+b)) + a)
	} else {
		fmt.Println(a*(n/(a+b)) + n%(a+b))
	}

}

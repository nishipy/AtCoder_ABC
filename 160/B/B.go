package main

import (
	"fmt"
)

func main() {
	var x int
	fmt.Scan(&x)

	fmt.Println((x/500)*1000 + ((x%500)/5)*5)
}

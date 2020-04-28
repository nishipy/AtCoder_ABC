package main

import "fmt"

func main() {
	var x, y, z string
	// &はアドレス
	fmt.Scan(&x, &y, &z)

	fmt.Println(z + " " + x + " " + y)
}

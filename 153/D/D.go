package main

import (
	"fmt"
	"math"
)

func main() {
	var h int
	fmt.Scan(&h)
	ans := 1
	i := 1
	fmt.Println(test(ans, h, i))

}

func test(r int, h int, i int) int {
	if h == 1 {
		return r
	} else {
		j := float64(i)
		r += int(math.Pow(2, j))
		h /= 2
		i++
		return test(r, h, i)
	}
}

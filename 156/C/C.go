package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextStr() string {
	sc.Scan()
	return sc.Text()
}

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()

	x := make([]int, n)
	sum := 0.0
	for i := 0; i < n; i++ {
		x[i] = nextInt()
		sum += float64(x[i])
	}

	p := int(sum/float64(n) + 0.5)
	ans := 0

	for _, y := range x {
		ans += (p - y) * (p - y)
	}

	fmt.Println(ans)
}

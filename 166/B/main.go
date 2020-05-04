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
	N, K := nextInt(), nextInt()
	A := make([]int, N)
	for i := 0; i < K; i++ {
		d := nextInt()
		for j := 0; j < d; j++ {
			l := nextInt()
			A[l-1]++
		}
	}

	ans := 0
	for _, a := range A {
		if a == 0 {
			ans++
		}
	}
	fmt.Println(ans)
}

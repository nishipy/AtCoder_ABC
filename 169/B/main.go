package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//Util
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

func maxInt(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func minInt(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func absInt(a int) int {
	if a >= 0 {
		return a
	} else {
		return -a
	}
}

func absFloat64(a float64) float64 {
	if a >= 0 {
		return a
	} else {
		return -a
	}
}

func minFloat64(a, b float64) float64 {
	if a > b {
		return b
	} else {
		return a
	}
}

func str2Int(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

//Main
var sc = bufio.NewScanner(os.Stdin)

const MAX = 1000000000000000000

var N int

func main() {
	sc.Split(bufio.ScanWords)
	N = nextInt()
	A := make([]string, N)
	for i := 0; i < N; i++ {
		a := nextStr()
		A[i] = a
		if a == "0" {
			fmt.Println(0)
			return
		}
	}

	//0がないとき
	var ans, b int64
	ans = 1
	b = MAX
	for _, a := range A {
		aa, _ := strconv.ParseInt(a, 10, 64)
		if aa > b {
			fmt.Println(-1)
			return
		}
		ans *= aa
		b /= aa
	}

	fmt.Println(ans)

}

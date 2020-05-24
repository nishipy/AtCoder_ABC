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
var q int
var X, Y string

func lcs(X, Y string) int {
	lenX := len(X)
	lenY := len(Y)
	dp := make([][]int, lenX+1)
	for i := 0; i < lenX+1; i++ {
		dp[i] = make([]int, lenY+1)
	}

	res := 0
	for i := 1; i < lenX+1; i++ {
		for j := 1; j < lenY+1; j++ {
			if X[i-1] == Y[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = maxInt(dp[i-1][j], dp[i][j-1])
			}
			//fmt.Println(dp)
			res = maxInt(res, dp[i][j])
		}

	}
	return res
}

func main() {
	sc.Split(bufio.ScanWords)
	q = nextInt()

	for i := 0; i < q; i++ {
		X = nextStr()
		Y = nextStr()

		fmt.Println(lcs(X, Y))
	}
}

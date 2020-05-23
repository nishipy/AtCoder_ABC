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
var N, W int
var weight, value [110]int
var dp [110][10010]int

func main() {
	sc.Split(bufio.ScanWords)
	N, W = nextInt(), nextInt()

	for i := 0; i < N; i++ {
		value[i] = nextInt()
		weight[i] = nextInt()
	}

	//初期条件
	//dp[0][w] = 0
	for w := 0; w <= W; w++ {
		dp[0][w] = 0
	}

	//dpループ
	for i := 0; i < N; i++ {
		for w := 0; w <= W; w++ {
			if w >= weight[i] {
				//i番目の品物を選べる場合
				dp[i+1][w] = maxInt(dp[i][w-weight[i]]+value[i], dp[i][w])
			} else {
				//i番目の品物を選べない場合
				dp[i+1][w] = dp[i][w]
			}
		}
	}

	fmt.Println(dp[N][W])
}

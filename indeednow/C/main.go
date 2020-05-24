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
var N, M int

//dp[i][j][k]:= 能力i, j, k の人の最高給与
var dp [102][102][102]int

func main() {
	sc.Split(bufio.ScanWords)
	N, M = nextInt(), nextInt()

	for i := 0; i < N; i++ {
		x, y, z, w := nextInt(), nextInt(), nextInt(), nextInt()
		//同じx, y, zが入力されたら、大きい方を採用する
		dp[x][y][z] = maxInt(dp[x][y][z], w)

	}

	//dpテーブルの作成
	for x := 0; x <= 100; x++ {
		for y := 0; y <= 100; y++ {
			for z := 0; z <= 100; z++ {
				dp[x][y][z] = maxInt(dp[x][y][z], maxInt(dp[maxInt(0, x-1)][y][z], maxInt(dp[x][maxInt(0, y-1)][z], dp[x][y][maxInt(0, z-1)])))
			}
		}
	}

	for i := 0; i < M; i++ {
		x, y, z := nextInt(), nextInt(), nextInt()
		fmt.Println(dp[x][y][z])
	}
}

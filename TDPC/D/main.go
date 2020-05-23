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
var N, D int
var dp [100][100][50][50]float64

func main() {
	sc.Split(bufio.ScanWords)
	N, D = nextInt(), nextInt()

	n2, n3, n5 := 0, 0, 0
	for D%2 == 0 {
		D /= 2
		n2++
	}
	for D%3 == 0 {
		D /= 3
		n3++
	}
	for D%5 == 0 {
		D /= 5
		n5++
	}
	if D != 1 {
		fmt.Println(0.0)
		return
	}

	//dp[i][a][b][c]:
	//i回サイコロを振って、積を素因数分解すると2がa回、3がb回,5がc回現れる確率
	dp[0][0][0][0] = 1

	for i := 0; i < N; i++ {
		for a := 0; a <= n2; a++ {
			for b := 0; b <= n3; b++ {
				for c := 0; c <= n5; c++ {
					dp[i+1][a][b][c] += (1.0 / 6) * dp[i][a][b][c]
					dp[i+1][minInt(a+1, n2)][b][c] += (1.0 / 6) * dp[i][a][b][c]
					dp[i+1][a][minInt(b+1, n3)][c] += (1.0 / 6) * dp[i][a][b][c]
					dp[i+1][minInt(a+2, n2)][b][c] += (1.0 / 6) * dp[i][a][b][c]
					dp[i+1][a][b][minInt(c+1, n5)] += (1.0 / 6) * dp[i][a][b][c]
					dp[i+1][minInt(a+1, n2)][minInt(b+1, n3)][c] += (1.0 / 6) * dp[i][a][b][c]
				}
			}
		}
	}

	fmt.Println(dp[N][n2][n3][n5])
}

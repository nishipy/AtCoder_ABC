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

//Main
var sc = bufio.NewScanner(os.Stdin)
var n int
var A [][]int

func main() {
	sc.Split(bufio.ScanWords)
	n = nextInt()
	A = make([][]int, 2)
	for i := 0; i < n; i++ {
		tmp := nextInt()
		A[0] = append(A[0], tmp)
	}
	sum := 0
	for i := 0; i < n; i++ {
		tmp := nextInt()
		sum += tmp
		A[1] = append(A[1], tmp)
	}

	upsum := make([]int, n)
	downsum := make([]int, n)
	upsum[0] = A[0][0]
	downsum[0] = sum

	ans := -1
	for i := 0; i < n; i++ {
		if upsum[i]+downsum[i] > ans {

			ans = upsum[i] + downsum[i]
		}
		if i < n-1 {
			upsum[i+1] = upsum[i] + A[0][i+1]
			downsum[i+1] = downsum[i] - A[1][i]
		}
	}

	fmt.Println(ans)
}

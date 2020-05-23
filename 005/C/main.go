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
var T, N, M int

func main() {
	sc.Split(bufio.ScanWords)
	T, N = nextInt(), nextInt()
	A := make([]int, N)
	for i := 0; i < N; i++ {
		A[i] = nextInt()
	}

	M = nextInt()
	B := make([]int, M)
	for i := 0; i < M; i++ {
		B[i] = nextInt()
	}

	indx := 0
	for _, b := range B {
		if indx >= N {
			fmt.Println("no")
			return
		}
		for j := indx; j < N; j++ {
			//たこ焼きが買えない場合
			if A[j] > b || A[j] < b-T {
				if j == N-1 {
					fmt.Println("no")
					return
				}
				continue
			}
			//fmt.Println(A[j], b)
			indx = j + 1
			break
		}

	}

	fmt.Println("yes")

}

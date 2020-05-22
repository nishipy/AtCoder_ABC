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
var N int

func main() {
	sc.Split(bufio.ScanWords)
	N = 1000 - nextInt()
	ans := 0
	for N > 0 {
		//fmt.Println(N)
		if N >= 500 {
			ans++
			N -= 500
		} else if N >= 100 {
			ans++
			N -= 100
		} else if N >= 50 {
			ans++
			N -= 50
		} else if N >= 10 {
			ans++
			N -= 10
		} else if N >= 5 {
			ans++
			N -= 5
		} else if N >= 1 {
			ans++
			N--
		}
	}

	fmt.Println(ans)
}

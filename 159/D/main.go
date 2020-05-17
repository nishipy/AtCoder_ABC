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

//Main
var sc = bufio.NewScanner(os.Stdin)
var n int

func main() {
	sc.Split(bufio.ScanWords)
	n = nextInt()
	A := make([]int, n)
	ma := make(map[int]int, n)
	for i := 0; i < n; i++ {
		A[i] = nextInt()
		ma[A[i]]++
	}

	all := 0
	for _, v := range ma {
		all += v * (v - 1) / 2
	}
	//fmt.Println(ma)
	for k := 0; k < n; k++ {
		v := ma[A[k]]
		before := v * (v - 1) / 2
		after := (v - 2) * (v - 1) / 2
		fmt.Println(all - (before - after))
	}
}

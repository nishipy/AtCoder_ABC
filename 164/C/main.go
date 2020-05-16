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
var s string

func main() {
	sc.Split(bufio.ScanWords)
	n = nextInt()
	m := make(map[string]int, n)
	for i := 0; i < n; i++ {
		s = nextStr()
		if _, ok := m[s]; ok {
			m[s]++
		} else {
			m[s] = 1
		}
	}
	//fmt.Println(m)
	fmt.Println(len(m))
}

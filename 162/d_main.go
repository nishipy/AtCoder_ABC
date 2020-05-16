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
var r, g, b int

func main() {
	sc.Split(bufio.ScanWords)
	n = nextInt()
	s = nextStr()

	for i := 0; i < n; i++ {
		switch string(s[i]) {
		case "R":
			r++
		case "B":
			b++
		case "G":
			g++
		}
	}
	ans := r * g * b
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			//j-i!=k-jの反例を除きたい
			k := j*2 - i
			//色が被る条件は、ansの初期値を定義した時点で除かれているのでスキップ
			if k >= n || s[i] == s[j] || s[k] == s[i] || s[k] == s[j] {
				continue
			}

			//j-i!=k-jの反例を除く
			ans--
		}
	}

	fmt.Println(ans)
}

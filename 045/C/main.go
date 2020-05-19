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
var s string

func main() {
	sc.Split(bufio.ScanWords)
	s = nextStr()
	n := len(s) - 1
	sum := 0
	//組み合わせの数

	for i := 0; i < (1 << uint(n)); i++ {
		w := []rune{}
		for l := 0; l <= n; l++ {
			//1が出るまで、wに文字列をためていく
			w = append(w, rune(s[l]))
			//パターンiについて、l番目のしきりが有効であるとき
			if (i & (1 << uint(l))) > 0 {
				//ここまでの文字列を数字として足す
				sum += str2Int(string(w))
				//wを空にする
				w = []rune{}
			}
		}
		//残ったwを足す
		sum += str2Int(string(w))
		//fmt.Println(sum)
	}

	fmt.Println(sum)
}

package main

import (
	"bufio"
	"fmt"
	"math"
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

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func pow(p, q int) int {
	return int(math.Pow(float64(p), float64(q)))
}

//Main
const (
	initialBufSize = 10000
	maxBufSize     = 1000000
	mod            = 1e9 + 7
)

var (
	sc *bufio.Scanner = func() *bufio.Scanner {
		sc := bufio.NewScanner(os.Stdin)
		buf := make([]byte, initialBufSize)
		sc.Buffer(buf, maxBufSize)
		return sc
	}()
)

var s, c string
var q, t, f int

func main() {
	sc.Split(bufio.ScanWords)
	s = nextStr()
	q = nextInt()

	head := make([]byte, 0, pow(10, 6))
	tail := make([]byte, 0, pow(10, 6))

	//反転する回数
	//奇数なら逆順になっている
	r := 0
	for i := 0; i < q; i++ {
		t = nextInt()
		if t == 1 {
			r++
		} else {
			f, c = nextInt(), nextStr()
			if f == 1 {
				if r%2 == 0 {
					head = append(head, c[0])
				} else {
					tail = append(tail, c[0])
				}
			} else {
				if r%2 == 0 {
					tail = append(tail, c[0])
				} else {
					head = append(head, c[0])
				}
			}
		}
	}

	if r%2 == 1 {
		s = reverse(s)
		s = reverse(string(tail)) + s + string(head)
	} else {
		s = reverse(string(head)) + s + string(tail)
	}

	fmt.Println(s)
}

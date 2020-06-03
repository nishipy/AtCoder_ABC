package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
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

func maxFloat64(a, b float64) float64 {
	if a > b {
		return a
	} else {
		return b
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

var n int
var s string
var patterns []string

func permute(a []string, l, r int) {
	if l == r {
		//fmt.Println(a)
		patterns = append(patterns, strings.Join(a, ""))
	} else {
		for i := l; i <= r; i++ {
			a[l], a[i] = a[i], a[l]
			permute(a, l+1, r)
			a[l], a[i] = a[i], a[l] //backtrack
		}
	}
}

func main() {
	sc.Split(bufio.ScanWords)
	n = nextInt()

	var p, q string
	for i := 1; i <= n; i++ {
		p += nextStr()
	}
	for i := 1; i <= n; i++ {
		q += nextStr()
	}

	nums := make([]string, n)
	for i := 1; i <= n; i++ {
		//fmt.Println(strconv.Itoa(i))
		nums[i-1] = string(strconv.Itoa(i))
	}

	a, b := -1, -1
	permute(nums, 0, len(nums)-1)
	//fmt.Println(patterns)
	//fmt.Println(p, q)
	sort.Sort(SortBy(patterns))
	//fmt.Println(patterns)
	for i, v := range patterns {
		if a > -1 && b > -1 {
			break
		}
		if reflect.DeepEqual(p, v) {
			a = i
		}
		if reflect.DeepEqual(q, v) {
			b = i
		}
	}

	fmt.Println(absInt(a - b))

}

type SortBy []string

func (a SortBy) Len() int           { return len(a) }
func (a SortBy) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortBy) Less(i, j int) bool { return a[i] < a[j] }

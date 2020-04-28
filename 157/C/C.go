package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

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

func main() {
	sc.Split(bufio.ScanWords)
	var n, m int
	n = nextInt()
	m = nextInt()

	ma := map[int]int{}

	if m == 0 {
		if n == 1 {
			ma[0] = 0
		} else {
			ma[0] = 1
		}
	} else {
		for l := 0; l < m; l++ {
			var s, c int
			s = nextInt() - 1
			c = nextInt()
			_, ok := ma[s]
			if ok && ma[s] != c {
				fmt.Println(-1)
				return
			} else if n > 1 && s == 0 && c == 0 {
				fmt.Println(-1)
				return
			} else {
				ma[s] = c
			}
		}
	}

	ans := 0
	for i := 0; i < n; i++ {
		_, ok := ma[i]
		if !ok {
			if i == 0 && n > 1 {
				ma[i] = 1
			} else {
				ma[i] = 0
			}
		}
		ans += ma[i] * int(math.Pow10(n-i-1))
	}
	fmt.Println(ans)

}

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

func PrimeFactors(n int) (pfs []int) {
	// Get the number of 2s that divide n
	for n%2 == 0 {
		pfs = append(pfs, 2)
		n = n / 2
	}

	// n must be odd at this point. so we can skip one element
	// (note i = i + 2)
	for i := 3; i*i <= n; i = i + 2 {
		// while i divides n, append i and divide n
		for n%i == 0 {
			pfs = append(pfs, i)
			n = n / i
		}
	}

	// This condition is to handle the case when n is a prime number
	// greater than 2
	if n > 2 {
		pfs = append(pfs, n)
	}

	return
}

//Main
var sc = bufio.NewScanner(os.Stdin)

const MAX = 1000000000000000000

var N int

func main() {
	N = nextInt()
	pfs := PrimeFactors(N)
	m := make(map[int]int, len(pfs))
	for i := 0; i < len(pfs); i++ {
		s := pfs[i]
		//mapにkeyが存在するか判定
		if _, ok := m[s]; ok {
			m[s]++
		} else {
			m[s] = 1
		}
	}

	ans := 0
	//fmt.Println(m)
	for _, v := range m {
		u := 1
		for {
			v -= u
			if v < 0 {
				break
			}
			ans++
			u++

			//fmt.Println(u, q, v, ans)
		}
	}

	fmt.Println(ans)
}

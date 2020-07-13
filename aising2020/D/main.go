package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
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

func powInt(p, q int) int {
	return int(math.Pow(float64(p), float64(q)))
}

func isPrime(x int) bool {
	if x == 1 {
		return false
	}
	if x == 2 {
		return true
	}
	if x%2 == 0 {
		return false
	}

	b := true
	rootx := int(math.Sqrt(float64(x)))
	i := 3
	for i <= rootx {
		if x%i == 0 {
			b = false
			break
		}
		i += 2
	}
	return b
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

func PrimeFactorsMap(n int) map[int]int {
	pfs := map[int]int{}
	// Get the number of 2s that divide n
	for n%2 == 0 {
		if _, ok := pfs[2]; ok {
			pfs[2]++
		} else {
			pfs[2] = 1
		}
		//pfs = append(pfs, 2)
		n = n / 2
	}

	// n must be odd at this point. so we can skip one element
	// (note i = i + 2)
	for i := 3; i*i <= n; i = i + 2 {
		// while i divides n, append i and divide n
		for n%i == 0 {
			if _, ok := pfs[i]; ok {
				pfs[i]++
			} else {
				pfs[i] = 1
			}
			n = n / i
		}
	}

	// This condition is to handle the case when n is a prime number
	// greater than 2
	if n > 2 {
		if _, ok := pfs[n]; ok {
			pfs[n]++
		} else {
			pfs[n] = 1
		}
	}

	return pfs
}

func sumInts(x []int) int {

	total := 0
	for _, v := range x {
		total += v
	}

	return total
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(x, y int) int {

	return x * y / gcd(x, y)

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

func main() {
	sc.Split(bufio.ScanWords)
	N, S := nextInt(), nextStr()
	//もとのpopcount
	pc := strings.Count(S, "1")

	//数字の文字列から、[]intへの変換
	X := make([]int, N)
	for i := 0; i < N; i++ {
		X[i] = int(S[i] - '0')
	}

	ans := make([]int, N)

	//Xの各桁は当然1または0
	//1の桁と0の桁に対して処理を行う
	for b := 0; b < 2; b++ {
		npc := pc
		if b == 0 {
			npc++
		} else {
			npc--
		}
		if npc <= 0 {
			continue
		}

		//もとの数の余り
		r0 := 0
		for i := 0; i < N; i++ {
			r0 = (r0 * 2) % npc
			r0 += X[i]
		}

		//各桁を反転した余り
		k := 1
		for i := N - 1; i >= 0; i-- {
			//Xの桁がb(0 or 1)と一致した場合のみ答えを求める
			if X[i] == b {
				r := r0
				if b == 0 {
					r = (r + k) % npc
				} else {
					//()内が負になるのを防ぐため、npcを足す
					r = (r - k + npc) % npc
				}
				ans[i] = f(r) + 1
			}
			k = (k * 2) % npc
		}
	}

	buf := bufio.NewWriter(os.Stdout)
	for i := 0; i < N-1; i++ {
		buf.Write([]byte(fmt.Sprintf("%d\n", ans[i])))
	}
	buf.Write([]byte(fmt.Sprintf("%d", ans[N-1])))
	buf.Flush()
}

//popcount: 2進数表記したときの1の個数を返す
func popcount(x int) int {
	c := 0
	for x != 0 {
		c += x & 1
		x >>= 1
	}
	return c
}

func f(x int) int {
	if x == 0 {
		return 0
	}
	return f(x%popcount(x)) + 1
}

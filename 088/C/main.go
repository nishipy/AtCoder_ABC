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

//Main
var sc = bufio.NewScanner(os.Stdin)
var c [][]int

func main() {
	sc.Split(bufio.ScanWords)
	c = make([][]int, 3)
	for i := 0; i < 3; i++ {
		c[i] = []int{nextInt(), nextInt(), nextInt()}
	}

	ans := "No"

	for a1 := 0; a1 <= 100; a1++ {
		b1 := c[0][0] - a1
		b2 := c[0][1] - a1
		b3 := c[0][2] - a1
		if b1 < 0 || b2 < 0 || b3 < 0 {
			continue
		}
		if c[1][0]-b1 != c[1][1]-b2 || c[1][1]-b2 != c[1][2]-b3 {
			continue
		}
		if c[2][0]-b1 != c[2][1]-b2 || c[2][1]-b2 != c[2][2]-b3 {
			continue
		}
		ans = "Yes"
		break
	}

	fmt.Println(ans)

}

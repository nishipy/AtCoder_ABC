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
var S []string
var h, w int

func main() {
	sc.Split(bufio.ScanWords)
	h, w = nextInt(), nextInt()
	S = make([]string, h)

	for i := 0; i < h; i++ {
		S[i] = nextStr()
	}

	l := []int{1, 0, -1, 0}
	m := []int{0, 1, 0, -1}

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if string(S[i][j]) == "#" {
				for k := 0; k < 4; k++ {
					x := i + l[k]
					y := j + m[k]
					if x >= 0 && x < h && y < w && y >= 0 {
						if string(S[x][y]) == "#" {
							break
						} else {
							if k == 3 {
								fmt.Println("No")
								return
							}
						}
					}
				}
			}
		}
	}

	fmt.Println("Yes")
}

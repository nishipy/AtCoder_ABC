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
var bombs []Bomb

type Bomb struct {
	i, j int
}

func main() {
	sc.Split(bufio.ScanWords)
	h, w = nextInt(), nextInt()
	S = make([]string, h)
	ans := make([][]int, h)

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			ans[i] = append(ans[i], 0)
		}
	}

	for i := 0; i < h; i++ {
		S[i] = nextStr()
		for j := 0; j < w; j++ {

			if string(S[i][j]) == "#" {
				tmp := Bomb{i, j}
				bombs = append(bombs, tmp)
				ans[i][j] = -2
			}

		}
	}

	for _, bomb := range bombs {
		i := bomb.i
		j := bomb.j
		for a := maxInt(i-1, 0); a <= minInt(i+1, h-1); a++ {
			for b := maxInt(j-1, 0); b <= minInt(j+1, w-1); b++ {
				// fmt.Println(i, j, a, b)
				if i == a && j == b {
					continue
				}
				if ans[a][b] == -2 {
					continue
				}
				ans[a][b]++

			}
		}
	}

	// fmt.Println(ans, bombs)
	buf := bufio.NewWriter(os.Stdout)
	for _, v := range ans {
		for i := 0; i < w; i++ {
			if v[i] == -2 {
				buf.Write([]byte(fmt.Sprintf("#")))
			} else {
				buf.Write([]byte(fmt.Sprintf("%d", v[i])))
			}
			if i == w-1 {
				buf.Write([]byte(fmt.Sprintf("\n")))
			}
		}
	}
	buf.Flush()

}

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
var h, w int
var c []string
var s, g Point

type Point struct {
	j, i int
}

func main() {
	sc.Split(bufio.ScanWords)
	h, w = nextInt(), nextInt()
	c = make([]string, h)
	done := make([][]int, h)
	for i := 0; i < h; i++ {
		c[i] = nextStr()
		for j, v := range string(c[i]) {
			done[i] = append(done[i], 0)
			if string(v) == "g" {
				g.i = j
				g.j = i
			} else if string(v) == "s" {
				s.i = j
				s.j = i
			}
		}
	}

	dx := []int{-1, 0, 0, 1}
	dy := []int{0, 1, -1, 0}

	stack := []Point{s}
	for len(stack) > 0 {
		sta := len(stack) - 1
		target := stack[sta]
		stack = stack[:sta]
		//fmt.Println(target)
		done[target.j][target.i] = 1
		if target.i == g.i && target.j == g.j {
			fmt.Println("Yes")
			return
		}
		for i := 0; i < 4; i++ {
			visiti := target.i + dx[i]
			visitj := target.j + dy[i]
			if visiti < 0 || visiti >= w || visitj < 0 || visitj >= h {
				continue
			}
			// fmt.Println(target.i, target.j, visiti, visitj, h, w)
			// fmt.Println(string(c[visiti][visitj]), done)
			// fmt.Println(done[visiti][visitj])
			// fmt.Println(string(c[visiti][visitj]))
			if done[visitj][visiti] == 0 && string(c[visitj][visiti]) != "#" {
				stack = append(stack, Point{visitj, visiti})
			}
		}
	}
	fmt.Println("No")

}

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
var R, C int
var s, g Point
var c []string

type Point struct {
	y, x, prey, prex int
}

func main() {
	sc.Split(bufio.ScanWords)
	R, C = nextInt(), nextInt()
	s.y, s.x = nextInt()-1, nextInt()-1
	g.y, g.x = nextInt()-1, nextInt()-1

	c = make([]string, R)
	done := make([][]int, R)
	for i := 0; i < R; i++ {
		c[i] = nextStr()
		for j := 0; j < C; j++ {
			done[i] = append(done[i], -1)
		}
	}
	//左、上、下、右移動
	dy := []int{0, 1, -1, 0}
	dx := []int{-1, 0, 0, 1}

	s.prex, s.prey = 0, 0
	q := []Point{s}
	done[s.y][s.x] = 0
	//fmt.Println(c, done, s, g)
	for len(q) > 0 {
		//キューから取り出し
		u := q[0]
		q = q[1:]

		//未訪問なら、スタートからの距離を入れる
		if s.y == u.y && s.x == u.x {
			done[u.y][u.x] = 0
		} else if done[u.y][u.x] == -1 {
			done[u.y][u.x] = done[u.prey][u.prex] + 1
		} else {
			continue
		}

		if u.x == g.x && u.y == g.y {
			fmt.Println(done[u.y][u.x])
			return
		}

		//次の訪問先をキューに入れる
		for i := 0; i < 4; i++ {
			nexty := u.y + dy[i]
			nextx := u.x + dx[i]
			//fmt.Println(nexty, nextx, done[nexty][nextx])
			if nexty < 0 || nexty >= R || nextx < 0 || nextx >= C {
				continue
			}
			//訪問済みであればスキップ
			if done[nexty][nextx] >= 0 {
				for i := 0; i < R; i++ {
					//fmt.Println(done[i])
				}
				continue
			}

			if done[nexty][nextx] == -1 && string(c[nexty][nextx]) != "#" {
				q = append(q, Point{nexty, nextx, u.y, u.x})
				//fmt.Println(q)
			}
		}
	}

}

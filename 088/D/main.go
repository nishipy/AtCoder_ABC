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
var H, W int
var dist [110][110]int

type Point struct {
	y, x int
}

func main() {
	sc.Split(bufio.ScanWords)
	H, W = nextInt(), nextInt()

	c := make([]string, H)
	whnum := 0
	for i := 0; i < H; i++ {
		c[i] = nextStr()
		for j := 0; j < W; j++ {
			if string(c[i][j]) != "#" {
				whnum++
			}
			dist[i][j] = -1
		}
	}

	//左、上、下、右移動
	dy := []int{0, 1, -1, 0}
	dx := []int{-1, 0, 0, 1}

	q := []Point{}
	q = append(q, Point{0, 0})
	dist[0][0] = 1
	//fmt.Println(c, done, s, g)
	for len(q) > 0 {
		//キューから取り出し
		u := q[0]
		q = q[1:]

		//次の訪問先をキューについて
		for i := 0; i < 4; i++ {
			nexty := u.y + dy[i]
			nextx := u.x + dx[i]
			//枠外ならスキップ
			if nexty < 0 || nexty >= H || nextx < 0 || nextx >= W {
				continue
			}
			//未訪問かつ通れる場合は評価
			if dist[nexty][nextx] == -1 && string(c[nexty][nextx]) != "#" {
				dist[nexty][nextx] = dist[u.y][u.x] + 1
				q = append(q, Point{nexty, nextx})
				//fmt.Println(q)
			}
		}
	}

	if dist[H-1][W-1] == -1 {
		fmt.Println(-1)
	} else {
		fmt.Println(whnum - dist[H-1][W-1])
	}

}

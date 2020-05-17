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
var n, m int

const INFTY = 1 << 21

type Node struct {
	prev, dist int
}

func bfs(nodes []Node, e [][]int) {
	init := Node{-1, INFTY}
	for i := 1; i < n; i++ {
		nodes[i] = init
	}
	nodes[0] = Node{-1, 0}
	var q []int
	q = append(q, 0)
	for len(q) > 0 {
		v := q[0]
		//queueから取り出したノードに隣接するノードについて
		for _, u := range e[v] {
			if nodes[u].dist > nodes[v].dist+1 {
				nodes[u].prev = v
				nodes[u].dist = nodes[v].dist + 1
				q = append(q, u)
			}
		}
		q = q[1:]
	}
}

func main() {
	sc.Split(bufio.ScanWords)
	n, m = nextInt(), nextInt()

	//隣接するノードを表す配列
	e := make([][]int, n)
	for i := 0; i < m; i++ {
		u, v := nextInt(), nextInt()
		u--
		v--
		e[u] = append(e[u], v)
		e[v] = append(e[v], u)
	}

	nodes := make([]Node, n)
	bfs(nodes, e)
	//fmt.Println(nodes)

	for i := 0; i < n; i++ {
		if nodes[i].dist == INFTY {
			fmt.Println("No")
			return
		}
	}

	fmt.Println("Yes")
	for i := 1; i < n; i++ {
		fmt.Println(nodes[i].prev + 1)
	}
}

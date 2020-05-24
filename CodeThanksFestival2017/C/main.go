package main

import (
	"bufio"
	"container/heap"
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
var N, K int
var A, B []int

const MAX = 100002

func main() {
	sc.Split(bufio.ScanWords)
	N, K = nextInt(), nextInt()
	A, B = make([]int, N), make([]int, N)
	for i := 0; i < N; i++ {
		A[i] = nextInt()
		B[i] = nextInt()
	}

	//pqの生成
	tmp := make(FactoryPQ, N)
	pq := &tmp
	heap.Init(pq)
	for i := 0; i < N; i++ {
		obj := &Item{priority: A[i], index: i}
		heap.Push(pq, obj)
	}

	ans := 0
	for i := 0; i < K; i++ {
		popped := heap.Pop(pq).(*Item)
		ans += popped.priority

		//優先度を更新して、
		//所要時間が入るので、大きいほど、優先度は低い
		//fmt.Println(popped.index, popped)
		popped.priority += B[popped.index]
		heap.Push(pq, popped)
	}

	fmt.Println(ans)
}

//優先度付きキュー
//https://gist.github.com/bxcodec/3cfb499a484441c0d74051b30c7cd123
//https://golang.org/pkg/container/heap/

type Item struct {
	priority int
	index    int
}
type FactoryPQ []*Item

func (pq FactoryPQ) Len() int           { return len(pq) }
func (pq FactoryPQ) Less(i, j int) bool { return pq[i].priority < pq[j].priority } // <: ASC, >: DESC
func (pq FactoryPQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *FactoryPQ) Push(x interface{}) {
	item := x.(*Item)
	*pq = append(*pq, item)
}
func (pq *FactoryPQ) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

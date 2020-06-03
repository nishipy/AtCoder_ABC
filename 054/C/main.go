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
var N, M int
var matrix [8][]int
var ans int

func dfs(done []bool, now, depth int) {

	//再帰の終了条件
	if depth == N-1 {
		ans++
		//すべての節を訪問できたら、ansをインクリメントして
		//for文に戻り、次の候補を探索
		//ここでは①に隣接する次の候補
		return
	}

	//訪問中の点を、訪問済みにする
	done[now] = true

	//次の候補について
	for _, nv := range matrix[now] {
		//訪問済みの場合はスキップ
		if done[nv] {
			continue
		}
		//未訪問の場合は、再帰呼び出し
		dfs(done, nv, depth+1)
	}

	//次の候補が無くなったら、自分を未訪問とする
	//最初に呼び出した時の、上のfor文に戻り、次の候補を探索
	done[now] = false

}

func main() {
	sc.Split(bufio.ScanWords)
	N, M = nextInt(), nextInt()

	for i := 0; i < M; i++ {
		var a, b int
		a, b = nextInt(), nextInt()
		a--
		b--
		matrix[a] = append(matrix[a], b)
		matrix[b] = append(matrix[b], a)
	}

	done := make([]bool, N)
	ans = 0
	dfs(done, 0, 0)

	fmt.Println(ans)
}

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

func dfs(v int, done []int) {
	//訪問中の点の他に、訪問していない点があれば、続ける
	end := true
	for i := 0; i < N; i++ {
		if done[i] != 1 && i != v {
			end = false
		}
	}
	if end {
		ans++
		return
	}

	//訪問中の点を、訪問済みにする
	done[v] = 1
	//次の候補について
	for _, nv := range matrix[v] {
		//訪問済みの場合はスキップ
		if done[nv] == 1 {
			continue
		}
		//未訪問の場合は、再帰呼び出し
		dfs(nv, done)
	}

	//次の候補が無くなったら、自分を未訪問として、おわり
	done[v] = 0
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

	done := make([]int, N)
	ans = 0
	dfs(0, done)

	fmt.Println(ans)
}

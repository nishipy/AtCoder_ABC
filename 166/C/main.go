package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

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

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func Min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func main() {
	sc.Split(bufio.ScanWords)
	N, M := nextInt(), nextInt()
	H := make([]int, N)
	maxs := make([]int, N)
	for i := 0; i < N; i++ {
		H[i] = nextInt()
	}
	for i := 0; i < M; i++ {
		a := nextInt() - 1
		b := nextInt() - 1
		maxs[a] = Max(maxs[a], H[b])
		maxs[b] = Max(H[a], maxs[b])
	}

	ans := 0

	//fmt.Println(maxs, H)
	for i := 0; i < N; i++ {
		if maxs[i] == 0 {
			ans++
			continue
		} else {
			if maxs[i] < H[i] {
				ans++
			}
		}
	}

	fmt.Println(ans)
}

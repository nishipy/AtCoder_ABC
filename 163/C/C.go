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

func main() {
	sc.Split(bufio.ScanWords)
	var n int
	n = nextInt()
	ma := make(map[int]int, n)
	for i := 0; i < n-1; i++ {
		tmp := nextInt()
		ma[tmp]++
	}

	for j := 1; j < n+1; j++ {
		fmt.Println(ma[j])
	}

}

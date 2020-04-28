package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	n := nextInt()
	a := make([]string, n)
	for i := range a {
		a[i] = nextStr()
	}

	sort.Strings(a)

	//最も多い回数を求める
	count, max := 1, 1
	for i := 1; i < n; i++ {
		if a[i] == a[i-1] {
			count++
			if max < count {
				max = count
			}
		} else {
			count = 1
		}
	}

	//最も多い回数と一致したら文字を出力
	if max == 1 {
		for i := range a {
			fmt.Println(a[i])
		}
	} else {
		ans := 1
		for i := 1; i < n; i++ {
			if a[i] == a[i-1] {
				ans++
			} else {
				ans = 1
			}
			if ans == max {
				fmt.Println(a[i-1])
			}
		}
	}
}

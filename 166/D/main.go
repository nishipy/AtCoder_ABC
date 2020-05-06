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
	X := nextInt()

	for i := -118; i < 120; i++ {
		for j := -119; j < 119; j++ {
			if (i*i*i*i*i)-(j*j*j*j*j) == X {
				fmt.Println(i, j)
				return
			}
		}
	}
}

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

//Main
var sc = bufio.NewScanner(os.Stdin)
var a, b, c, d int

func main() {
	sc.Split(bufio.ScanWords)
	a, b, c, d = nextInt(), nextInt(), nextInt(), nextInt()
	iTaka := 0
	iAoki := 0
	for c > 0 {
		c -= b
		iTaka += 1
	}

	for a > 0 {
		a -= d
		iAoki += 1
	}

	//fmt.Println(iTaka, iAoki)

	if iTaka <= iAoki {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

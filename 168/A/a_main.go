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
var n int

func main() {
	sc.Split(bufio.ScanWords)
	n = nextInt()
	test := n % 10
	if test == 3 {
		fmt.Println("bon")
	} else if test == 0 || test == 1 || test == 6 || test == 8 {
		fmt.Println("pon")
	} else {
		fmt.Println("hon")
	}
}

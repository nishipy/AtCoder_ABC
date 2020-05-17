package main

import (
	"bufio"
	"fmt"
	"math"
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
var a, b, h, m int

func main() {
	sc.Split(bufio.ScanWords)
	a, b, h, m = nextInt(), nextInt(), nextInt(), nextInt()
	minutes := 60*h + m
	var aAngle, bAngle float64
	//時針
	aAngle = 0.5 * float64(minutes%720)
	//分針
	bAngle = 6.0 * float64(minutes%60)

	tmpdiff := absFloat64(aAngle - bAngle)
	diff := minFloat64(tmpdiff, 360.0-tmpdiff)
	//fmt.Println(diff)

	c2 := float64(a*a+b*b) - 2*float64(a*b)*math.Cos(diff*math.Pi/180)
	fmt.Println(math.Sqrt(c2))

}

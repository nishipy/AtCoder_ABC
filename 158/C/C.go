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
	var a, b int
	a = nextInt()
	b = nextInt()

	leftA := float64(a) / 0.08
	rightA := float64(a+1) / 0.08
	leftB := float64(b) / 0.1
	rightB := float64(b+1)/0.1 - 1

	if a%2 != 0 {
		leftA++
		rightA--
	}

	//fmt.Println(leftA, leftB, rightA, rightB)
	ans := -1
	if leftA < leftB {
		for i := int(leftA); i <= int(rightA); i++ {
			if i == int(leftB) {
				ans = i
				break
			}
		}
	} else {
		for i := int(leftB); i <= int(rightB); i++ {
			if i == int(leftA) {
				ans = i
				break
			}
		}
	}

	fmt.Println(ans)

}

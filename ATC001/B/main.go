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

type unionTree struct {
	el     int
	parent *unionTree
}

func (u *unionTree) root() *unionTree {
	if u.parent == nil {
		return u
	}

	root := u.parent.root()
	//検索効率化
	if root != u.parent {
		u.parent = root
	}
	return root
}

func (u *unionTree) same(v *unionTree) bool {
	return u.root() == v.root()
}

func (u *unionTree) unite(v *unionTree) {
	uRoot := u.root()
	vRoot := v.root()
	if uRoot == vRoot {
		return
	}

	if uRoot.el < vRoot.el {
		vRoot.parent = uRoot
	} else {
		uRoot.parent = vRoot
	}
}

func main() {
	sc.Split(bufio.ScanWords)
	N, Q := nextInt(), nextInt()
	unions := make([]unionTree, N)
	for i := 0; i < N; i++ {
		unions[i] = unionTree{el: i}
	}

	for i := 0; i < Q; i++ {
		p, a, b := nextInt(), nextInt(), nextInt()
		if p == 0 {
			unions[a].unite(&unions[b])
		} else {
			if unions[a].same(&unions[b]) {
				fmt.Println("Yes")
			} else {
				fmt.Println("No")
			}
		}
	}
}

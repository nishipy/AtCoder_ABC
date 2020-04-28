package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	if n <= k {
		fmt.Println(0)
		return
	}

	h := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&h[i])
	}

	sort.Sort(sort.IntSlice(h))
	ans := 0
	for j := 0; j < n-k; j++ {
		ans += h[j]
	}
	fmt.Println(ans)
}

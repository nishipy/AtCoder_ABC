package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)

	s := map[string]int{}
	max := 1
	for i := 0; i < n; i++ {
		var tmp string
		fmt.Scan(&tmp)
		_, exists := s[tmp]
		if !exists {
			s[tmp] = 1
		} else {
			s[tmp]++
			if max < s[tmp] {
				max = s[tmp]
			}
		}
	}

	var ans []string
	for k, v := range s {
		if v == max {
			ans = append(ans, k)
		}
	}
	sort.Strings(ans)
	for _, a := range ans {
		fmt.Println(a)
	}

}

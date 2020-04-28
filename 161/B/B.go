package main

import "fmt"

func main() {
	var N, M int
	// コンソールで入力を要求
	fmt.Scan(&N, &M)

	t := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&t[i])
	}

	s := 0
	for _, x := range t {
		s += x
	}

	ans := 0
	for j := 0; j < N; j++ {
		if t[j]*4*M < s {
			ans++
		}
	}

	if ans+M <= N {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

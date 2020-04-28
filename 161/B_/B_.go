package B_

import (
	"fmt"
	"sort"
)

func B_() {
	var N, M, sum int
	fmt.Scan(&N, &M)
	a := make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Scan(&a[i])
		sum += a[i]
	}
	// 降順に並ぶ
	// * sort.IntSlice(a) -> 型を変換
	// * sort.Reverse -> 上のデータを降順に並び替え
	sort.Sort(sort.Reverse(sort.IntSlice(a)))

	ok := true

	for i := 0; i < M; i++ {
		if a[i]*4*M < sum {
			ok = false
		}
	}

	if ok {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}

package main

import (
	"fmt"
)

func main() {
	var k, n int
	fmt.Scan(&k, &n)
	nums := scanNums(n)

	dist := nums[0] + k - nums[n-1]

	for i := 0; i < n-1; i++ {
		diff := nums[i+1] - nums[i]
		if diff > dist {
			dist = diff
		}
	}

	fmt.Println(k - dist)
}

func scanNums(len int) (nums []int) {
	var num int
	for i := 0; i < len; i++ {
		fmt.Scan(&num)
		nums = append(nums, num)
	}
	return
}

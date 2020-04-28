package main

import (
	"fmt"
)

func main() {
	var S string
	fmt.Scan(&S)
	// s[i] は、i番目のバイトを表す
	// 一般的に、英数字は1文字1バイト
	if S[2] == S[3] && S[4] == S[5] {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)

	//文字列の繰り返しには、strings.Repeat()を使う
	fmt.Println(strings.Repeat("x", len(s)))

}

# AtCoder_ABC
My solutions for AtCoder Beginner Contest problems.

## Tips - Golang
### Input data
```go
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
	sc.Split(bufio.ScanWords) //スペース区切り
	// sc.Split(bufio.ScanLines) //1行まるごと
	var n int
	n = nextInt()
...
```

### Min(), Max() for int
```go
func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func Min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
```

### Split a string on whitespace
```go
l := strings.Split(str, " ")
```


### Join array to a string with whitespace
```go
strings.Join(arr, ",")
```
### Cast String <-> Int
```go
var i int
var s string="123"
i, _ = strconv.Atoi(s)
fmt.Println(i) // -> 123
```
```go
var i int=321
var s string
s = strconv.Itoa(i)
fmt.Println(s) // -> "123"
```
```go
func str2Int(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}
```

### Display []int like as string
```go
//with buf.Flush()
func printList(list []int) {
	buf := bufio.NewWriter(os.Stdout)
	for _, v := range list {
		buf.Write([]byte(fmt.Sprintf(" %d", v)))
	}
	buf.Write([]byte("\n"))
	buf.Flush()
}

//Or
func arrayAsString(arr []int) string {
	return strings.TrimRight(fmt.Sprintf("%+v", arr)[1:], "]")
}

//Or
for i:= 0; i<len(array); i++ {
	fmt.Printf("%d ", array[i])
}
```

### 2進数に変換
```go
b := fmt.Sprintf("%b", 2)
//0埋め4文字
//b := fmt.Sprintf("%04b", 2)
fmt.Println(b)
```

### GCD
```go
func gcd(a, b int) int {
    if b == 0 {
        return a
    }
    return gcd(b, a%b)
}
```

### GCD with math/big
```go
func gcd(m, n uint64) uint64 {
    x := new(big.Int)
    y := new(big.Int)
    z := new(big.Int)
    a := new(big.Int).SetUint64(m)
    b := new(big.Int).SetUint64(n)
    return z.GCD(x, y, a, b).Uint64()
}
```

### Define and use map
```go
func main() {
	sc.Split(bufio.ScanWords)
	n = nextInt()
	m := make(map[string]int, n)
	for i := 0; i < n; i++ {
		s = nextStr()
		//mapにkeyが存在するか判定
		if _, ok := m[s]; ok {
			m[s]++
		} else {
			m[s] = 1
		}
	}
	//fmt.Println(m)
	fmt.Println(len(m))
}
```

### rune
* 文字列など、文字単位で扱う時に使う
* [String と　Rune](https://text.baldanders.info/golang/string-and-rune/)
  * 例
	```golang
	package main

	import "fmt"

	func main() {
		nihongo := "日本語"
		nihongoRune := []rune(nihongo)
		size := len(nihongoRune)

		fmt.Printf("nihongo = %d characters : ", size)
		for i := 0; i < size; i++ {
			fmt.Printf("%#U ", nihongoRune[i])
		}
		fmt.Printf("\n")
	}
	//-> nihongo = 3 characters : U+65E5 '日' U+672C '本' U+8A9E '語' 
	```
  * for range 文を使ったループでは、rune単位で取得される
  * string -> rune -> stringと変換していけば、安全に処理できる
* 例ABC061 C

### Sort
```go
// Intをsort(昇順）
sort.Ints(i)
fmt.Println(i)

// stringをアルファベット順にsort
sort.Strings(s)
fmt.Println(s)

// structをStringでsort (名前順)
sort.Slice(p, func(i, j int) bool { return p[i].Name < p[j].Name })
fmt.Println(p)

// structをIntでsort (昇順）
sort.Slice(p, func(i, j int) bool { return p[i].Point < p[j].Point })
fmt.Println(p)
```
または、
```go
sort.Sort(SortBy(cad))

type SortBy []string

func (a SortBy) Len() int           { return len(a) }
func (a SortBy) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortBy) Less(i, j int) bool { return a[i] < a[j] }

```


## アルゴリズム系[*1]
### グラフ問題
* 可視化サイト
  * https://hello-world-494ec.firebaseapp.com/
* 幅優先探索
  * 最短距離問題に使う
    * 各節点に隣接する点を入力する
    * 次に訪問する節点を保持するキューを用意する
      * 始点をキューに入れる
      * キューから取り出し、その点について処理する
      * 処理した節点に隣接する点を、キューに入れる
      * キューが空になるまで処理を繰り返す
    * 例
      * [168 D - .. (Double Dots)](./168/D/main.go)
* 深さ優先探索
  * 塗り潰し問題、全探索問題
  * 次の訪問先をスタックして、進めていく
  * 例
    * [ATC001 - A](./ATC001/A/main.go)
    * [ABC054 - C](./054/C/main.go)


### bit全探索
* 2^n通りの組み合わせを、2進数を用いて解くやつ
  * [Golangのビット演算](https://hydrocul.github.io/wiki/programming_languages_diff/number/bit-operator.html)
    * `&`, `|` : AND, OR
    * `<<`: 左シフト。整数を2倍
      * `1 << n`で、2^nを表す。
    * `>>`: 右シフト。整数を1/2倍
* 例
  * [045 C](./045/C/main.go)
  * [079 C](./079/C/main.go)

### Greedy
* [グリーディ法](http://www2.kobe-u.ac.jp/~ky/da2/haihu04.pdf)
  * ある段階で、最も利益の大きい部分解を選択していく
  * 必ずしも、最適解になるとは限らない
    * 高速なので、他の手法と組み合わせて使う
* 区間スケジューリング
  * 区間の終端もしくは始端でソートする
  * 例
    * [B - Robot Arms](./keyence2020/B/main.go)
* 双対問題
  * 実際は区間スケジューリング問題
  * 例
    * [103 D - Islands War](./103/D/main.go)
      * 終端でソートする
      * 共通部分ができる限りはOK
      * https://drken1215.hatenablog.com/entry/2018/07/21/224200
* 辞書式順序最小を求める問題
  * 不定のところには、'a'をいれて、候補のリストを作る
  * 候補をsortする
    * `sort.Strings(cad)`など
  * 例
    * [C - Dubious Document 2](./076/main.go)

## 動的計画法(DP)
* 値を覚えて再利用することで、処理を効率化する
* [ナップザック問題](https://onlinejudge.u-aizu.ac.jp/problems/DPL_1_B)
  * [解説記事](https://qiita.com/drken/items/a5e6fe22863b7992efdb)
    * `dp[i][w]`を、i番目までの商品の中から容量を超えないように選んだときの価値の総和の最大値とする
    * これを漸化式と見立てて解いていく
  	  ![解説記事から引用](https://qiita-user-contents.imgix.net/https%3A%2F%2Fqiita-image-store.s3.amazonaws.com%2F0%2F182963%2Ffcd3c29a-9f3e-7984-3549-21fa113fab26.jpeg?ixlib=rb-1.2.2&auto=format&gif-q=60&q=75&s=e12b1a6b61eb9cc97d783b848bdea6c4)
  * 例
    * [TDPC - A](TDPC/A/main.go)
    * [TDPC - C](TDPC/C/main.go)




---
[*1]: https://qiita.com/drken/items/e77685614f3c6bf86f44

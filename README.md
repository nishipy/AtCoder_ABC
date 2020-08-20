# kyopuro
My solutions for kyopuro contests problems.

## Tips - Golang
### make関数
* 引数
  * 左から、型、長さ、キャパシティー
  * キャパシティを省略した場合、長さと同じだけ確保する
```go
a :=  make([]int, 0, 5)
```
### Pointer
* `*`間接参照演算子
  * ポインタを定義。また、ポインタの中身を取得
* `&`: アドレス演算子
  * アドレスを取得
```go
package main

import (
	"fmt"
)

func main() {
	var pointer *int
	var n int = 100

	//nのアドレスを代入
	pointer = &n

	fmt.Println(&n, pointer) //nのアドレス
	fmt.Println(*pointer) //pointerの中身=nの値を表示
}
```
### Input data
* 注意‼︎
  * [Goのbufio.Scannerは入力データの1行の長さが一定以上になるとスキャンをやめてしまう](https://mickey24.hatenablog.com/entry/bufio_scanner_line_length)
  * [bufio.Scannerのエラーと解決策 #golang](https://inaba.hatenablog.com/entry/2017/02/19/045844)
    * 入力が大きい時は、単純に`fmt.Scan()`をする方がいい
    * もしくは、最大サイズを指定する

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	initialBufSize = 10000
	maxBufSize     = 1000000
	mod            = 1e9 + 7
)

var (
	sc *bufio.Scanner = func() *bufio.Scanner {
		sc := bufio.NewScanner(os.Stdin)
		buf := make([]byte, initialBufSize)
		sc.Buffer(buf, maxBufSize)
		return sc
	}()
)

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

### アルファベットの扱い
```go
fmt.Println('a') #-> 97
fmt.Println(string(97)) #-> a

S := nextStr()
al := make([]int, 26)
	for i := 0; i < N; i++ {
		al[S[i]-'a']++
	}
```

### int64でセイウチ
```go
ans := int64(0)
```

### 型ごとの最大値最小値
[`math`](https://golang.org/pkg/math/)を使う。

```go
import math

ans1 := math.MaxInt64
ans2 := math.MinInt64
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

### 2進数表記のときの1の数を求める
```go
func popcount(x int) int {
	c := 0
	for x != 0 {
		c += x & 1
		//1ビットずらす
		x >>= 1
	}
	return c
}
```

### 2進数に変換
```go
b := fmt.Sprintf("%b", 2)
//0埋め4文字
//b := fmt.Sprintf("%04b", 2)
fmt.Println(b)
```

### 数字の文字列を[]intに変換
```go
	X := make([]int, N)
	for i := 0; i < N; i++ {
		X[i] = int(S[i] - '0')
	}
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

### LCM
```go
//gcdを使う
func lcm(x, y int) int {

	return x * y / gcd(x, y)

}
```

### GCD/LCMのint64バージョン
```go
func nextInt64() int64 {
	var i int64
	if sc.Scan() {
		if num, err := strconv.ParseInt(sc.Text(), 10, 64); err == nil {
			i = num
		}
	}
	return i
}

func gcd64(a, b int64) int64 {
	if a%b != 0 {
		return gcd64(b, a%b)
	} else {
		return b
	}
}

func lcm64(a, b int64) int64 {
	return a / gcd64(a, b) * b
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
* Goでは、シングルクオートがrune型を表す
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

### 文字列を含むか
* [参考](https://qiita.com/tchnkmr/items/b3d0b884db8d7d91fb1b#%E6%96%87%E5%AD%97%E5%88%97%E4%B8%AD%E3%81%AB%E6%8C%87%E5%AE%9A%E6%96%87%E5%AD%97%E5%88%97%E3%81%8C%E5%90%AB%E3%81%BE%E3%82%8C%E3%81%A6%E3%81%84%E3%82%8B%E3%81%8B)
```go
s = "Alfa Bravo Charlie Delta Echo Foxtrot Golf"

//文字列中に指定文字列が含まれているか
fmt.Println(strings.Contains(s, "Delta")) // -> true

//出現位置
fmt.Println(strings.Index(s, "Delta")) // -> 19
fmt.Println(strings.Index(s, "Hotel")) // -> -1
fmt.Println(strings.LastIndex(s, "a")) // -> 23

//出現回数
fmt.Println(strings.Count(s, "a")) // -> 4
```
### 文字列の置き換え
* [参考](https://qiita.com/Sekky0905/items/f0bed43ad3ab4be13385)

### 文字列の反転
```go
func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
```

### 文字列の連結
* サイズが大きくなると、`[]byte`にappendするのが速い
  * [Golang の文字列連結はどちらが速い？](https://qiita.com/spiegel-im-spiegel/items/16ab7dabbd0749281227)
  *  例題
     *  [158 D](158/D/main.go)

### 文字列の中に含まれる、任意の文字列の位置
* https://ashitani.jp/golangtips/tips_string.html#string_Find
* [151 - A](151/A/main.go)
```
func main() {
	sc.Split(bufio.ScanWords)
	c = nextStr()

	alphabet := "abcdefghijklmnopqrstuvwxyz"

	i := strings.Index(alphabet, c)
	fmt.Println(string(alphabet[i+1]))

}
```

### 文字列の中に含まれる、任意の文字列の個数
* `strings.Count()`を使う
* 例題
  * [150 - B](./159/B/B.go)
```go
var patterns []string

func permute(a []string, l, r int) {
	if l == r {
		//fmt.Println(a)
		patterns = append(patterns, strings.Join(a, ""))
	} else {
		for i := l; i <= r; i++ {
			a[l], a[i] = a[i], a[l]
			permute(a, l+1, r)
			a[l], a[i] = a[i], a[l] //backtrack
		}
	}
}
```

### 累乗
* math.Pow()は、float64型なので、Int型を用意
```
func pow(p, q int) int {
	return int(math.Pow(float64(p), float64(q)))
}
```

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

* 複数要素によるソート
  * Aの値が同じ場合は、Bの値でソートするなど
  * 実装例: [128 - B](/128/B/main.go)
```go
type Catalog struct {
	name  string
	point int
	id    int
}

type SortBy []Catalog

func (a SortBy) Len() int      { return len(a) }
func (a SortBy) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a SortBy) Less(i, j int) bool {
	if a[i].name < a[j].name {
		return true
	} else if a[i].name == a[j].name {
		return a[i].point > a[j].point
	} else {
		return false
	}

```

#### string型のsort
```go
type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
    return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
    return len(s)
}

func SortString(s string) string {
    r := []rune(s)
    sort.Sort(sortRunes(r))
    return string(r)
}
```

### 素数判定
```
func isPrime(x int) bool {
	if x == 1 {
		return false
	}
	if x == 2 {
		return true
	}
	if x%2 == 0 {
		return false
	}

	b := true
	rootx := int(math.Sqrt(float64(x)))
	i := 3
	for i <= rootx {
		if x%i == 0 {
			b = false
			break
		}
		i += 2
	}
	return b
}
```

### 素因数分解
* 例題
  * [169 - D](./169/D/main.go)
```
func PrimeFactors(n int) (pfs []int) {
	// Get the number of 2s that divide n
	for n%2 == 0 {
		pfs = append(pfs, 2)
		n = n / 2
	}

	// n must be odd at this point. so we can skip one element
	// (note i = i + 2)
	for i := 3; i*i <= n; i = i + 2 {
		// while i divides n, append i and divide n
		for n%i == 0 {
			pfs = append(pfs, i)
			n = n / i
		}
	}

	// This condition is to handle the case when n is a prime number
	// greater than 2
	if n > 2 {
		pfs = append(pfs, n)
	}

	return
}
```

* Mapバージョン
```go
func PrimeFactorsMap(n int) map[int]int {
	pfs := map[int]int{}
	// Get the number of 2s that divide n
	for n%2 == 0 {
		if _, ok := pfs[2]; ok {
			pfs[2]++
		} else {
			pfs[2] = 1
		}
		//pfs = append(pfs, 2)
		n = n / 2
	}

	// n must be odd at this point. so we can skip one element
	// (note i = i + 2)
	for i := 3; i*i <= n; i = i + 2 {
		// while i divides n, append i and divide n
		for n%i == 0 {
			if _, ok := pfs[i]; ok {
				pfs[i]++
			} else {
				pfs[i] = 1
			}
			n = n / i
		}
	}

	// This condition is to handle the case when n is a prime number
	// greater than 2
	if n > 2 {
		if _, ok := pfs[n]; ok {
			pfs[n]++
		} else {
			pfs[n] = 1
		}
	}

	return pfs
}
```

### 約数列挙
* `a * a <= N`の範囲の、aを調べる
  * つまり、`Sqrt(N)`まで調べればいい
* Nがaで割り切れる時、約数である
* 例題
  * [144 - C](144/C/main.go)

### 約数の個数
* 1からNまでの数について、約数の個数を数え上げる場合
  * 長さNの配列Aを作る。`A[i]`の値が約数の個数となる
  * i = 1として、`A[(iの倍数)]`をインクリメントする
    * iの倍数がNを超えたら止める
* 例題
  * [172 - D](172/D/main.go)

### [エラトステネスの篩](https://ja.wikipedia.org/wiki/%E3%82%A8%E3%83%A9%E3%83%88%E3%82%B9%E3%83%86%E3%83%8D%E3%82%B9%E3%81%AE%E7%AF%A9)
* 指定された整数以下の全ての素数を発見するための単純なアルゴリズム
* 類題
  * [170 - D](170/D/main.go)

### 数列の増減
* 株価とか標高差とかが苦手な気がする
* 教訓
  * 数列の問題は、後ろからも考察する
  * 増加や減少が連続する数を考えてみる
* 例題
  * [AGC040 - A](AGC/040/A/main.go)
    * 左に連続する`<`の個数と右の連続する`>`の個数を考える

### 期待値の問題
* 確率pで成功する試行を、成功するまで繰り返したとき、繰り返す回数の期待値は`1/p`である
  * 例: [078 C - HSI](https://atcoder.jp/contests/abc078/tasks/arc085_a)
  * [参考サイト](https://drken1215.hatenablog.com/entry/2019/03/23/175300)
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
* 深さ優先探索(DFS)
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

```go
for bit := 0; bit < (1 << uint64(n)); bit++ {
  // fmt.Println(bit)
  // 3を入力した場合、0~7の計8つが出力される。
  for i := 0; i < n; i++ {
	// fmt.Println(i)
	// 3を入力した場合、0, 1, 2のセットが計8つが出力される。
    if (bit>>uint64(i))&1 == 1 { // bitsのi個目の要素の状態がonかどうかチェック(ここは問題によって条件を変化させる)
    }
  }
}
```

#### 整数を2進数に変換
* [参考サイト](https://ashitani.jp/golangtips/tips_num.html)
```go
package main

import "fmt"

func main() {
    s := ""
    s = fmt.Sprintf("%b", 255)
	fmt.Println(s) // => "11111111"
}
```
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
  * 他の例題
    * [TDPC - A](TDPC/A/main.go)
    * [TDPC - C](TDPC/C/main.go)
* [最長共通部分列][http://judge.u-aizu.ac.jp/onlinejudge/description.jsp?id=ALDS1_10_C&lang=jp]
  * コードは、[こちら](./AOJ/ALDS1_10_C/main.go#l75)
  * 文字列S1とS2の最長共通部分を求める問題
    * `dp[i][j]`に、`S1`のi文字目までと`S2`のj文字目までの最長共通部分列の長さを格納する
      * X[i] == Y[j]ならば、dp[i-1][j-1]+1を格納
        * i-1文字目とj-1文字目までのLCSの長さに1を足せばよい
  	* X[i] != Y[j]ならば、dp[i-1][j]とdp[i][j-1]の大きい方を格納
  * その他の例題
    * [C - Optimal Recommendations](https://atcoder.jp/contests/indeednow-finala-open/tasks/indeednow_2015_finala_c)

## Union-Find木
* グループ分けを、木構造で管理する
* [例題](https://atcoder.jp/contests/atc001/tasks/unionfind_a)
  * [実装例](./ATC001/B/main.go)

## データ構造
* priority_queue
  * [例題](https://atcoder.jp/contests/code-thanks-festival-2017-open/tasks/code_thanks_festival_2017_c)
  * [解答](./CodeThanksFestival2017/C/main.go)
  * 優先度付きキュー
  * キューに対して、要素を優先度付きで追加する
  * 最も優先度の高い要素を取り除き、それを返す




---
[*1]: https://qiita.com/drken/items/e77685614f3c6bf86f44

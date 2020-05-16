package main

import "testing"

//testing.T構造体
// テストを失敗させたり、テストメッセージの出力をするために使用
func Test_main(t *testing.T) {
	type args struct {
		k int
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{7, 500, 600}, want: "OK"},
		{name: "2", args: args{4, 5, 7}, want: "NG"},
		{name: "3", args: args{1, 11, 11}, want: "OK"},
	}
	for _, tt := range tests {
		testCase := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := solution(testCase.args.a, testCase.args.b, testCase.args.k); got != testCase.want {
				t.Fatalf("want %v, got %v", testCase.want, got)
			}
		})
	}
}

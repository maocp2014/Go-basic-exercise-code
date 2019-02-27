package main

import "testing"

// 表格驱动测试
func TestSubstr(t *testing.T) {
	tests := []struct {
		s   string
		ans int
	}{
		// Normal cases
		{"abcabcbb", 3},
		{"pwwkew", 3},

		// Edge cases
		{"", 0},
		{"b", 1},
		{"bbbbbbbbb", 1},
		{"abcabcabcd", 4},

		// chinese support
		{"我爱中国长城", 6},
		{"一二三三二一", 3},
		{"黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花", 8},
	}

	for _, tt := range tests{
		actual := lengthOfNonRepeatingSubStr_2(tt.s)
		if actual != tt.ans{
			t.Errorf("got %d for input %s; expected %d", actual, tt.s, tt.ans)
		}
	}
}

// // 性能测试
func BenchmarkSubstr(b *testing.B){
	s := "黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花"
	ans := 8

	for i := 0; i < b.N; i++{
		actual := lengthOfNonRepeatingSubStr_2(s)
		if actual != ans{
			b.Errorf("got %d for input %s; expected %d", actual, s, ans)
		}
	}
}

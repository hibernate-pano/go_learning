package operator_test

import "testing"

func TestCompareArray(t *testing.T) {
	// 初始化数组
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 3, 4, 5}
	//c := [...]int{1, 2, 3, 4, 5}
	d := [...]int{1, 2, 3, 4}
	e := [...]int{2, 1, 3, 4}
	// 数组比较
	t.Log(a == b)
	t.Log(a == d)
	t.Log(b == d)
	// 数组相等的条件是，位置和值都要保持一致
	t.Log(a == e)
}

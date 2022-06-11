package operator_test

import "testing"

func TestCompareArray(t *testing.T) {
	// 初始化数组
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 3, 4, 5}
	//c := [...]int{1, 2, 3, 4, 5}
	d := [...]int{1, 2, 3, 4}
	// 数组比较
	t.Log(a == b)
	t.Log(a == d)
	t.Log(b == d)
}

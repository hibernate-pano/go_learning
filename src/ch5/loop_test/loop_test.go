package loop_test

import "testing"

// 测试 while 循环
func TestWhileLoop(t *testing.T) {
	n := 0
	for n < 5 {
		n++
		t.Log(n)
	}
}

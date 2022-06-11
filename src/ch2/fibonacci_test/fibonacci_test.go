package fibonacci_test

import (
	"testing"
)

func TestFibonacciList(t *testing.T) {
	a, b := 1, 1
	t.Log(a)
	for i := 0; i < 5; i++ {
		t.Log(b)
		tmp := a
		a = b
		b = tmp + a
	}
}

// go 中可以不通过临时变量来进行数值交换
func TestExchangeValue(t *testing.T) {
	a, b := 1, 2
	a, b = b, a
	t.Log(a, b)
}

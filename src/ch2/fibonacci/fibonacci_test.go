package fibonacci

import (
	"fmt"
	"testing"
)

func TestFibonacciList(t *testing.T) {
	a, b := 1, 1
	fmt.Print(a + b)
	for i := 0; i < 5; i++ {
		tmp := a
		a = b
		b = tmp + a
	}
}

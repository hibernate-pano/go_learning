package type_test

import (
	"math"
	"testing"
)

type MyInt int64

// go 不支持隐式类型转换
func TestImplicit(t *testing.T) {
	var a int32 = 1
	var b int64
	b = int64(a)
	var c MyInt
	c = MyInt(b)
	t.Log(a, b, c)
}

func TestMath(t *testing.T) {
	a := math.MaxInt64
	b := math.MaxFloat64
	c := math.MaxUint32
	t.Log(a, b, c)
}

// go 不支持指针运算
func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	t.Log(a, aPtr)
	t.Logf("%T %T", a, aPtr)
}

func TestString(t *testing.T) {
	var s string
	t.Log("*" + s + "*˚")
	t.Log(len(s))
	if s == "" {
		t.Log("It`s Blank")
	}
}

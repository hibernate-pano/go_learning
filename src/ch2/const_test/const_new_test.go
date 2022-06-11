package const_test

import "testing"

// iota 可以进行常量的初始化赋值操作
// 递增
const (
	Monday = iota + 1
	Tuesday
	Wednesday
	ThursDay
	Friday
	Saturday
	Sunday
)

// 位运算
const (
	Readable = iota << 1
	writeable
	Executable
)

func TestConst(t *testing.T) {
	// 测试递增
	t.Log(Monday, Tuesday, Wednesday, ThursDay, Friday, Saturday, Sunday)
	// 测试位运算
	a := 1
	t.Log(a&Readable == Readable, a&writeable == writeable, a&Executable == Executable)
}

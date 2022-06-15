package interface_test

import "testing"

// 定义接口
type Programmer interface {
	// 待实现接口
	WriteCodes() string
}

// 定义结构体
type GoProgrammer struct {
}

// 实现方法
func (gg *GoProgrammer) WriteCodes() string {
	return "Game Over"
}
func TestProgrammerInterface(t *testing.T) {
	var gg Programmer
	gg = new(GoProgrammer)
	codes := gg.WriteCodes()
	t.Log(codes)
}

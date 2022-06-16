package interface_test

import "testing"

// 定义接口
type Programmer interface {
	// 待实现接口
	WriteCodes() string
	eatShit() string
}

// 定义结构体
type GoProgrammer struct {
}

func (gg *GoProgrammer) eatShit() string {
	return "Bia Ji Bia Ji"
}

// 实现方法
func (gg *GoProgrammer) WriteCodes() string {
	return "Game Over"
}

func TestProgrammerInterface(t *testing.T) {
	var gg Programmer = new(GoProgrammer)
	codes := gg.WriteCodes()
	sounds := gg.eatShit()
	t.Log(codes, sounds)
}

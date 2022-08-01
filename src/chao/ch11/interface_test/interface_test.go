package interface_test

import "testing"

// 定义接口
type Programmer interface {
	// 待实现接口
	WriteCodes(ee string) string
	eatShit(ee string) string
}

// 定义结构体
type GoProgrammer struct {
}

func (gg *GoProgrammer) eatShit(ee string) string {
	//return "Bia Ji Bia Ji"
	return ee
}

// 实现方法
func (gg *GoProgrammer) WriteCodes(ee string) string {
	//return "Game Over"
	return ee
}

func TestProgrammerInterface(t *testing.T) {
	var gg Programmer = new(GoProgrammer)
	codes := gg.WriteCodes("Game Over")
	sounds := gg.eatShit("Bia Ji Bia Ji")
	t.Log(codes, sounds)
}

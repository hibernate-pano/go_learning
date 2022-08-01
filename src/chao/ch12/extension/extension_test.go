package extension

import (
	"fmt"
	"testing"
)

// 定义 Pet 类型
type Pet struct {
}

// 定义 Pet 方法
func (p *Pet) Speak() {
	fmt.Println("...")
}

// 定义 Pe t方法
func (p *Pet) SpeakTo(name string) {
	p.Speak()
	fmt.Println("", name)
}

type Dog struct {
	p *Pet
}

// 定义 Dog 方法
func (d *Dog) Speak() {
	d.p.Speak()
}

// 定义 Dog 方法
func (d *Dog) SpeakTo(name string) {
	d.p.SpeakTo(name)
}

func TestInherit(t *testing.T) {
	dog := new(Dog)
	dog.SpeakTo("pano")
}

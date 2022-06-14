package map_ext

import (
	"math/rand"
	"testing"
)

// map 的值可以是一个函数
func TestMapWithFuncValue(t *testing.T) {
	m := map[int]func(param int) int{}
	// 初始化每个key对应的函数
	m[1] = func(param int) int {
		return param
	}
	m[2] = func(param int) int {
		return param * param
	}
	m[3] = func(param int) int {
		return param * param * param
	}
	// 测试
	t.Log(m[1](2), m[2](2), m[3](2))
}

// 通过map方法实现set数据结构
func TestMapForSet(t *testing.T) {
	// 初始化map
	m := map[int]int{}
	// 添加参数
	arr := [...]int{1, 2, 3, 4, 4, 4, 4, 5, 6, 1}
	// 返回去重的切片
	var res []int

	for _, v := range arr {
		// 判断key是否存在
		if _, b := m[v]; b {
			// 存在直接输出日志
			t.Logf("%d is Existing!", v)
		} else {
			// 不存在添加到切片，并更新该下标的值
			res = append(res, v)
			m[v] = rand.Int()
		}
	}
	// 输出
	t.Log(res)
}

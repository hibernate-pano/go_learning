package func_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestFuncMultiReturn(t *testing.T) {
	a, b := returnMultiValues()
	t.Log(a, b)
}

// 多返回值函数
func returnMultiValues() (int, int) {
	return rand.Int(), rand.Int()
}

// 创造一个函数，用以计算内部函数的执行时长
func timeSpent(inner func(param int) int) func(param int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent : ", time.Since(start).Seconds())
		return ret
	}
}

func slowFunc(param int) int {
	time.Sleep(time.Second * 5)
	return param
}

func TestFuncInFunc(t *testing.T) {
	timeFunc := timeSpent(slowFunc)
	timeFunc(10)
}

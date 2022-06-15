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

// 测试可变长度参数
func TestVarParam(t *testing.T) {
	t.Log(SumInt(1, 2, 3))
}

// 可变长度参数方法
func SumInt(params ...int) int {
	sum := 0
	for _, p := range params {
		sum += p
	}
	return sum
}

// 延迟执行defer
func TestDeferFunc(t *testing.T) {
	// defer 修饰的方法在方法返回之前或者抛出异常之前执行
	// 类似 java 的 try..catch..finally...结构中的finally，必定会执行的方法
	defer func() {
		t.Log("defer func executed!!")
	}()
	t.Log("begin")
	// 类似手动抛出异常信息，panic之后的方法不再执行
	panic("Fatal Err")
}

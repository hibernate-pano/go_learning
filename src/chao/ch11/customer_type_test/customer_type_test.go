package customer_type_test

import (
	"fmt"
	"time"
)

// 借用之前写好的函数
// 创造一个函数，用以计算内部函数的执行时长
func timeSpent(inner IntConv) IntConv {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent : ", time.Since(start).Seconds())
		return ret
	}
}

// 创建一个自定义func类型，或者理解成别名也行
type IntConv func(param int) int

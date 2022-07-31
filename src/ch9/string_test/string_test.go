package string_test

import (
	"strconv"
	"strings"
	"testing"
)

func TestStringFunc(t *testing.T) {
	s := "a,b,c"
	parts := strings.Split(s, ",")
	t.Log(parts)

	for _, part := range parts {
		t.Log(part)
	}

	ss := strings.Join(parts, "-")
	t.Log(ss)
}

func TestStringConv(t *testing.T) {
	// 数字转字符串
	s := strconv.Itoa(10)
	t.Log("string" + s)
	// 字符串转数字,需要显示处理转换错误逻辑
	if i, err := strconv.Atoi("20x"); err == nil {
		t.Log(10 + i)
	}

	atoi, err := strconv.Atoi("20x")
	if err != nil {
		return
	}
	t.Log(atoi)
}

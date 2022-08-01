package slice_test

import "testing"

func TestSliceInit(t *testing.T) {
	var s0 []int //slice 和 array 相似，但是不需要声明数组长度
	t.Log(len(s0), cap(s0))

	s0 = append(s0, 1)
	t.Log(len(s0), cap(s0))

	s1 := []int{1, 2, 3, 4, 5}
	t.Log(len(s1), cap(s1))

	// make 关键字
	s2 := make([]int, 3, 8)
	t.Log(len(s2), cap(s2))
	t.Log(s2[0], s2[1], s2[2])

	for i := 0; i < 100; i++ {
		s2 = append(s2, i)
		// slice的cap会自动扩容，每次扩容x2
		t.Log(len(s2), cap(s2))
	}
}

func TestSliceShareMemory(t *testing.T) {
	year := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12"}
	// 数组或者切片都可以使用这样的方式截取
	Q2 := year[3:6]
	t.Log(Q2, len(Q2), cap(Q2))

	summer := year[5:9]
	t.Log(summer, len(summer), cap(summer))

	// 因为使用的是一片连续的内存空间，假如 summer 变更了某一个内存地址的值，那么使用同一个slice的内存空间的所有的值都会变更。
	summer[0] = "Unknown"
	t.Log(summer, len(summer), cap(summer))
	t.Log(Q2, len(Q2), cap(Q2))
	t.Log(year, len(year), cap(year))
}

func TestSliceComparing(t *testing.T) {
	year1 := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12"}
	year2 := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12"}
	// slice 不可比较， 只能判断空或者nil
	//t.Log(year1 == year2)
	if year1 == nil {

	}

	if len(year2) > 0 {

	}
}

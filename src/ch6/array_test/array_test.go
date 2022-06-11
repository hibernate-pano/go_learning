package array_test

import "testing"

func TestArrayInit(t *testing.T) {
	var arr [3]int
	t.Log(arr[1], arr[2], arr[0])

	arr1 := [3]int{1, 2, 3}
	arr2 := [...]int{1, 2, 3, 4}
	t.Log(arr1, arr2)
}

func TestArrayLoop(t *testing.T) {
	arr2 := [...]int{1, 2, 3, 4}

	for i := 0; i < len(arr2); i++ {
		t.Log(arr2[i])
	}

	for i, e := range arr2 {
		t.Log(i, e)
	}

	for _, e := range arr2 {
		t.Log(e)
	}
}

func TestArraySubstring(t *testing.T) {
	a := [...]int{1, 2, 3, 4, 5}
	t.Log(a[:3])
	t.Log(a[1:3])
	t.Log(a[1:])
}

package condition_test

import "testing"

func TestIfMultiSec(t *testing.T) {
	if _, err := someFun(); err == nil {
		t.Log("no err")
	} else {
		t.Log("err")
	}
}

func someFun() (int, error) {
	return 0, nil
}

// 测试switch多case的情况
func TestSwitchMultiCase(t *testing.T) {
	for i := 0; i < 10; i++ {
		switch i {
		case 0, 2, 4, 6:
			t.Log("Even")
		case 1, 3, 5, 7:
			t.Log("Odd")
		default:
			t.Log("Out Of Limit")
		}
	}
}

package map_test

import "testing"

// map 只有len 没有cap
func TestInitMap(t *testing.T) {
	m := map[string]int{"one": 1, "two": 2}
	t.Logf("len m  %d", len(m))
	m1 := map[string]int{}
	m1["one"] = 1
	m1["two"] = 2
	t.Logf("len m1 %d", len(m1))
	m2 := make(map[string]int, 10)
	m2["key"] = 1
	t.Logf("len m2 %d", len(m2))
	t.Log(m, m1, m2)
}

func TestAccessNotExistingKey(t *testing.T) {
	m := map[string]int{"one": 1, "two": 2}
	// 访问不存在的key，会默认为零值
	t.Log(m["three"])
	m["three"] = 3
	// 手动判断零值还是值=0
	if v, ok := m["three"]; ok {
		t.Logf("Existing! value = %d", v)
	} else {
		t.Log("Not Existing!")
	}
}

func TestMapRange(t *testing.T) {
	m := map[string]int{"one": 1, "two": 2}
	m["three"] = 3
	for k, v := range m {
		t.Log(k, v)
	}
}

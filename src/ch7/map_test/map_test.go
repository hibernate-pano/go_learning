package map_test

import "testing"

func TestMap(t *testing.T) {
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

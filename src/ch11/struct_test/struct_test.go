package struct_test

import "testing"

// 定义类型 struct 结构体
type Employee struct {
	Id   string
	Name string
	Age  int
}

func TestCreateEmployeeEntity(t *testing.T) {
	// 张三
	e := Employee{
		Id:   "0",
		Name: "张三",
		Age:  18,
	}
	e.Name = "张大山"
	// 李四
	e1 := Employee{"1", "李四", 30}
	// 王五
	e2 := new(Employee) // 返回指针
	e2.Id = "2"
	e2.Name = "王五"
	e2.Age = 27

	a := 10
	t.Logf("e is : %T", e)
	t.Logf("e is : %X", &e)
	t.Logf("e is : %X", &e.Id)
	t.Logf("e is : %X", &e.Name)
	t.Logf("e1 is : %T", &e1)
	t.Logf("a location is : %X", &a)
	t.Logf("e2 is : %T", e2)
	t.Logf("e2 is : %X", &e2)
	t.Logf("e2 is : %X", &e2.Id)
}

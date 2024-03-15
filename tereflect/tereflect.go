package tereflect

import "fmt"

type Studentd struct {
	Name  string
	Age   int
	Name2 string
}

func (stu Studentd) SetName(name string, name2 string) {

	stu.Name = name
	stu.Name2 = name2
}

func (stu Studentd) SetAge(age int) {
	stu.Age = age
	fmt.Println(stu.Age)
}

func (stu Studentd) Print() string {
	return fmt.Sprintf("Name:%s,age: %d, name2: %s", stu.Name, stu.Age, stu.Name2)
}

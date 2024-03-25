package model

import "fmt"

type student struct {
	name string
	age  int
}

// 当一个结构体是私有的只能在本包访问，可以通过工厂模式让其他包也能获取这个结构体实例和它的字段
func (student *student) String() string {
	return fmt.Sprintf("student.name=%s, student.age=%d", student.name, student.age)
}

func (student *student) GetStudentName() string {
	return student.name
}

func (student *student) GetStudentAge() int {
	return student.age
}

func NewStudent(name string, age int) *student {
	return &student{
		name: name,
		age:  age,
	}
}

package model

import "fmt"

// 将person和person的部分字段进行封装，即只能在本包使用
type person struct {
	Name string
	age  int
	sal  float64
}

func NewPerson(name string) *person {
	return &person{Name: name}
}

func (p *person) SetAge(age int) {
	if age > 0 && age < 150 {
		p.age = age
	} else {
		fmt.Println("age error")
	}
}

func (p *person) GetAge() int {
	return p.age
}

func (p *person) SetSal(sal float64) {
	if sal > 0 && sal < 300000 {
		p.sal = sal
	} else {
		fmt.Println("sal error")
	}
}

func (p *person) GetSal() float64 {
	return p.sal
}

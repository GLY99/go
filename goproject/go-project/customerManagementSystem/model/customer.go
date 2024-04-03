package model

import "fmt"

type Customer struct {
	Id     int
	Name   string
	Gender string
	Age    int
	Phone  string
	Email  string
}

func NewCustom(id int, name string, gender string, age int,
	phone string, email string) *Customer {
	return &Customer{
		Id: id, Name: name, Gender: gender, Age: age, Phone: phone,
		Email: email,
	}
}

func NewCustomNoId(name string, gender string, age int,
	phone string, email string) *Customer {
	return &Customer{
		Name: name, Gender: gender, Age: age, Phone: phone,
		Email: email,
	}
}

func (c *Customer) ShowCustomInfo() string {
	return fmt.Sprintf("%d\t%s\t%s\t%d\t%s\t%s", c.Id, c.Name, c.Gender,
		c.Age, c.Phone, c.Email)
}

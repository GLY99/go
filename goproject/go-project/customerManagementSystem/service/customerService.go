package service

import (
	"cms/model"
	"fmt"
)

type CustomerService struct {
	customers   []*model.Customer
	customNumer int
}

func NewCustomService() *CustomerService {
	return &CustomerService{
		customers:   make([]*model.Customer, 0, 500),
		customNumer: 1,
	}
}

func (cs *CustomerService) ListCustomers() []*model.Customer {
	return cs.customers
}

func (cs *CustomerService) AddCustomer(customer *model.Customer) {
	customer.Id = cs.customNumer
	cs.customers = append(cs.customers, customer)
	cs.customNumer++
	fmt.Println("添加客户成功")
}

func (cs *CustomerService) FindCustomerById(id int) int {
	for idx, customer := range cs.customers {
		if customer.Id == id {
			return idx
		}
	}
	return -1
}

func (cs *CustomerService) DeleteCustomer(id int) {
	idx := cs.FindCustomerById(id)
	if idx == -1 {
		fmt.Printf("id为%d的客户不存在!\n", id)
	} else {
		cs.customers = append(cs.customers[:idx], cs.customers[idx+1:]...)
		fmt.Printf("删除id为%d的客户成功\n", id)
	}
}

func (cs *CustomerService) UpdateCustomer(id int, name string, gender string, age int,
	phone string, email string) {
	idx := cs.FindCustomerById(id)
	if idx == -1 {
		fmt.Printf("id为%d的客户不存在!\n", id)
	} else {
		customer := cs.customers[idx]
		if name != "" {
			customer.Name = name
		}
		if gender != "" {
			customer.Gender = gender
		}
		if age != -1 {
			customer.Age = age
		}
		if phone != "" {
			customer.Phone = phone
		}
		if email != "" {
			customer.Email = email
		}
		fmt.Printf("修改id为%d的客户信息成功\n", id)
	}
}

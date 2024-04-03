package view

import (
	"cms/model"
	"cms/service"
	"fmt"
)

var maxRetryCount = 3

type customView struct {
	input           int
	loop            bool
	customerService *service.CustomerService
}

func NewCustomView() *customView {
	return &customView{loop: true, customerService: service.NewCustomService()}
}

func (cv *customView) exitSystem() {
	fmt.Printf("你确定要退出吗?(Y/N):")
	var flag byte
	for {
		fmt.Scanf("%c\n", &flag)
		if flag == 'Y' {
			cv.loop = false
			return
		} else if flag == 'N' {
			fmt.Println("请继续操作")
			return
		} else {
			if maxRetryCount > 1 {
				fmt.Printf("选择错误,你还有%d次机会,请输入Y/N:", maxRetryCount-1)
			}
		}
		maxRetryCount--
		if maxRetryCount <= 0 {
			break
		}
	}
	fmt.Println("退出系统失败!")
}

func (cv *customView) list() {
	customers := cv.customerService.ListCustomers()
	fmt.Println("--------------------客户列表--------------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for _, customer := range customers {
		fmt.Println(customer.ShowCustomInfo())
	}
	fmt.Println("------------------------------------------------")
}

func (cv *customView) add() {
	var name, gender, phone, email string
	var age int
	fmt.Printf("请输入客户的姓名:")
	fmt.Scanf("%s\n", &name)
	fmt.Printf("请输入客户的性别:")
	fmt.Scanf("%s\n", &gender)
	fmt.Printf("请输入客户的年龄:")
	fmt.Scanf("%d\n", &age)
	fmt.Printf("请输入客户的电话:")
	fmt.Scanf("%s\n", &phone)
	fmt.Printf("请输入客户的邮箱:")
	fmt.Scanf("%s\n", &email)
	customer := model.NewCustomNoId(name, gender, age, phone, email)
	cv.customerService.AddCustomer(customer)
}

func (cv *customView) delete() {
	fmt.Printf("请输入要删除客户的编号:")
	var id int
	fmt.Scanf("%d\n", &id)
	fmt.Printf("确认要删除这个客户吗(Y/N)?")
	var flag byte
	for {
		fmt.Scanf("%c\n", &flag)
		if flag == 'Y' {
			cv.customerService.DeleteCustomer(id)
			return
		} else if flag == 'N' {
			fmt.Println("请继续操作")
			return
		} else {
			if maxRetryCount > 1 {
				fmt.Printf("选择错误,你还有%d次机会,请输入Y/N:", maxRetryCount-1)
			}
		}
		maxRetryCount--
		if maxRetryCount <= 0 {
			break
		}
	}
	fmt.Println("删除客户失败!")
}

func (cv *customView) update() {
	fmt.Printf("请输入要修改客户的编号:")
	var id int
	fmt.Scanf("%d\n", &id)
	var name, gender, phone, email string
	var age int = -1
	fmt.Printf("请输入修改后的客户的姓名(回车表示不修改):")
	fmt.Scanf("%s\n", &name)
	fmt.Printf("请输入修改后的客户的性别(回车表示不修改):")
	fmt.Scanf("%s\n", &gender)
	fmt.Printf("请输入修改后的客户的年龄(回车表示不修改):")
	fmt.Scanf("%d\n", &age)
	fmt.Printf("请输入修改后的客户的电话(回车表示不修改):")
	fmt.Scanf("%s\n", &phone)
	fmt.Printf("请输入修改后的客户的邮箱(回车表示不修改):")
	fmt.Scanf("%s\n", &email)
	fmt.Printf("确认要对这个客户进行修改吗(Y/N)?")
	var flag byte
	for {
		fmt.Scanf("%c\n", &flag)
		if flag == 'Y' {
			cv.customerService.UpdateCustomer(id, name, gender, age, phone, email)
			return
		} else if flag == 'N' {
			fmt.Println("请继续操作")
			return
		} else {
			if maxRetryCount > 1 {
				fmt.Printf("选择错误,你还有%d次机会,请输入Y/N:", maxRetryCount-1)
			}
		}
		maxRetryCount--
		if maxRetryCount <= 0 {
			break
		}
	}
	fmt.Println("删除客户失败!")
}

func (cv *customView) ShowMainMenu() {
	for {
		fmt.Println("--------------------客户信息管理软件--------------------")
		fmt.Println("                    1.添加客户")
		fmt.Println("                    2.修改客户")
		fmt.Println("                    3.删除客户")
		fmt.Println("                    4.客户列表")
		fmt.Println("                    5.退    出")
		fmt.Printf("                    请选择(1-5):")
		fmt.Scanf("%d\n", &cv.input)
		switch cv.input {
		case 1:
			cv.add()
		case 2:
			cv.update()
		case 3:
			cv.delete()
		case 4:
			cv.list()
		case 5:
			cv.exitSystem()
		default:
			fmt.Println("输入错误,请重新输入")
		}
		if !cv.loop {
			break
		}
	}
}

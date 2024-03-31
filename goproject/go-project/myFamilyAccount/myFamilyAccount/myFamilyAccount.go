package myFamilyAccount

import "fmt"

type familyAccount struct {
	user       string
	password   string
	userSelect int8
	balance    float64
	money      float64
	notes      string
	details    string
	flag       bool
	exit       bool
}

func NewFamilyAccount() *familyAccount {
	return &familyAccount{
		user: "xixi", password: "123456",
		userSelect: -1, balance: 0,
		money: 0, notes: "",
		details: "收支\t账户金额\t收支金额\t说明\n",
		flag:    false, exit: false,
	}
}

func (fa *familyAccount) auth() bool {
	fmt.Printf("请输入用户名:")
	var user string
	fmt.Scanf("%s\n", &user)
	fmt.Printf("请输入密码:")
	var password string
	fmt.Scanf("%s\n", &password)
	if fa.user == user && fa.password == password {
		return true
	}
	fmt.Println("用户名或密码错误!")
	return false
}

func (fa *familyAccount) showDetails() {
	if fa.flag {
		fmt.Println("-------------------当前收支明细-------------------")
		fmt.Println(fa.details)
	} else {
		fmt.Println("当前没有任何收支！")
	}
}

func (fa *familyAccount) income() {
	fmt.Println("-------------------登记收入-------------------")
	fmt.Printf("本次收入金额:")
	fmt.Scanf("%f\n", &fa.money)
	fmt.Printf("本次收入说明:")
	fmt.Scanf("%s\n", &fa.notes)
	fa.balance += fa.money
	fa.details += fmt.Sprintf("收入\t%.2f\t%.2f\t%s\n", fa.balance, fa.money, fa.notes)
	fa.flag = true
}

func (fa *familyAccount) pay() {
	fmt.Println("-------------------登记支出-------------------")
	fmt.Printf("本次支出金额:")
	fmt.Scanf("%f\n", &fa.money)
	fmt.Printf("本次支出说明:")
	fmt.Scanf("%s\n", &fa.notes)
	if fa.money > fa.balance {
		fmt.Println("账户余额不足!")
	} else {
		fa.balance -= fa.money
		fa.details += fmt.Sprintf("支出\t%.2f\t%.2f\t%s\n", fa.balance, fa.money, fa.notes)
		fa.flag = true
	}
}

func (fa *familyAccount) exitSys() {
	fmt.Printf("你确定要退出吗?Y/N:")
	var choice byte
	fmt.Scanf("%c\n", &choice)
	if choice == 'Y' {
		fmt.Println("退出软件")
		fa.exit = true
	} else if choice == 'N' {
		fmt.Println("请继续操作")
	} else {
		fmt.Println("你的输入错误请重新输入!")
	}
}

func (fa *familyAccount) ShowMainMenu() {
lable1:
	for {
		fmt.Println("-------------------家庭收支记账软件-------------------")
		fmt.Println("                   1.收支明细")
		fmt.Println("                   2.登记收入")
		fmt.Println("                   3.登记支出")
		fmt.Println("                   4.退出软件")
		fmt.Printf("                   请选择1-4:")
		fmt.Scanf("%d\n", &fa.userSelect)
		switch fa.userSelect {
		case 1:
			if fa.auth() {
				fa.showDetails()
			}
		case 2:
			if fa.auth() {
				fa.income()
			}
		case 3:
			if fa.auth() {
				fa.pay()
			}
		case 4:
			fa.exitSys()
			if fa.exit {
				break lable1
			}
		default:
			fmt.Println("错误的选择，请重新选择")
		}
	}
}

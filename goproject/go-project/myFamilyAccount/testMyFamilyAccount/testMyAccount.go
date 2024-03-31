package test

import "fmt"

func StartMyAccount() {
	var userSelect int8
	var balance float64
	var money float64
	var notes string
	var details string = "收支\t账户金额\t收支金额\t说明\n"
	var flag bool = false
lable1:
	for {
		fmt.Println("-------------------家庭收支记账软件-------------------")
		fmt.Println("                   1.收支明细")
		fmt.Println("                   2.登记收入")
		fmt.Println("                   3.登记支出")
		fmt.Println("                   4.退出软件")
		fmt.Printf("                   请选择1-4:")
		fmt.Scanf("%d\n", &userSelect)
		switch userSelect {
		case 1:
			if flag {
				fmt.Println("-------------------当前收支明细-------------------")
				fmt.Println(details)
			} else {
				fmt.Println("当前没有任何收支！")
			}
		case 2:
			fmt.Println("-------------------登记收入-------------------")
			fmt.Printf("本次收入金额:")
			fmt.Scanf("%f\n", &money)
			fmt.Printf("本次收入说明:")
			fmt.Scanf("%s\n", &notes)
			balance += money
			details += fmt.Sprintf("收入\t%.2f\t%.2f\t%s\n", balance, money, notes)
			flag = true
		case 3:
			fmt.Println("-------------------登记支出-------------------")
			fmt.Printf("本次支出金额:")
			fmt.Scanf("%f\n", &money)
			fmt.Printf("本次支出说明:")
			fmt.Scanf("%s\n", &notes)
			if money > balance {
				fmt.Println("账户余额不足!")
			} else {
				balance -= money
				details += fmt.Sprintf("支出\t%.2f\t%.2f\t%s\n", balance, money, notes)
				flag = true
			}
		case 4:
			fmt.Printf("你确定要退出吗?Y/N:")
			var choice byte
			fmt.Scanf("%c\n", &choice)
			if choice == 'Y' {
				fmt.Println("退出软件")
				break lable1
			} else if choice == 'N' {
				fmt.Println("请继续操作")
			} else {
				fmt.Println("你的输入错误请重新输入!")
			}
		default:
			fmt.Println("错误的选择，请重新选择")
		}
	}
}

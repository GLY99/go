package main

import (
	"myFamilyAccount/myFamilyAccount"
	// test "myFamilyAccount/testMyFamilyAccount"
)

func main() {
	// test.StartMyAccount()
	fa := myFamilyAccount.NewFamilyAccount()
	fa.ShowMainMenu()
}

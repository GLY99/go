package main

import (
	"fmt"
	"golangStudy/factoryModel/model"
)

func main() {
	student1 := model.NewStudent("xiaoming", 18)
	fmt.Println(student1)                  // student.Name=xiaoming, student.Age=18
	fmt.Println(student1.GetStudentName()) // xiaoming
	fmt.Println(student1.GetStudentAge())  // 18
}

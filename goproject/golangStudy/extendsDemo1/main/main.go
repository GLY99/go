package main

import "fmt"

type Student struct {
	Name  string
	Age   int
	Score float64
}

type Home struct {
	address string
}

// 公共方法
func (home *Home) showHomeInfo() {
	fmt.Printf("home address=%s\n", home.address)
}

func (student *Student) ShowInfo() {
	fmt.Printf("student name=%s, age=%d, score=%.2f\n", student.Name, student.Age, student.Score)
}

func (student *Student) SetScore(score float64) {
	student.Score = score
}

type Pupil struct {
	Student
	Home
	Hobbies string // 特有的字段
}

// 特有的一些方法
func (pupil *Pupil) showHomeInfo() {
	fmt.Printf("pupil home info not support show\n")
}

func (pupil *Pupil) testing() {
	fmt.Println("小学生在考试")
}

type Graduate struct {
	Student
	Home
	Hobbies string // 特有的字段
}

// 特有的一些方法
func (pupil *Graduate) testing() {
	fmt.Println("大学生在考试")
}

func main() {
	//  最基本的使用方式
	pupil1 := Pupil{}
	pupil1.Student.Name = "xiaoming"
	pupil1.Student.Age = 12
	pupil1.testing() // 小学生在考试
	pupil1.Student.SetScore(100)
	pupil1.Student.ShowInfo() // student name=xiaoming, age=12, score=100.00

	graduate1 := Graduate{}
	graduate1.Student.Name = "xiaohong"
	graduate1.Student.Age = 12
	graduate1.testing() // 大学生在考试
	graduate1.Student.SetScore(99)
	graduate1.Student.ShowInfo() // student name=xiaohong, age=12, score=99.00

	// 继承后的结构体初始化其他方式, 结构体名字可以作为字段
	pupil2 := Pupil{Home: Home{address: "address"}, Student: Student{Name: "haha", Age: 18}} // 初始化时带有字段可以随意变化顺序和缺省字段
	pupil3 := Pupil{Student{Name: "haha", Age: 18}, Home{address: "address"}, "hobbies"}     // 初始化时没有字段必须按照顺序初始化而且不能缺少字段
	// pupil4 := Pupil{Student{Name: "haha", Age: 18}} // 如果继承了多个结构体并且只初始化其中一部分不能省略结构体名称, 因为无法判断这个值给哪个结构体
	fmt.Println(pupil2, pupil3) // {{haha 18 0} {address} } {{haha 18 0} {address} hobbies}

	// 匿名结构体的小写字段和方法都可以被访问
	fmt.Println(pupil2.Home.address) // address
	pupil3.showHomeInfo()            // home address=address

	// 匿名结构体的字段访问可以简化,先寻找本结构体有没有Name,如果没有找匿名结构体有没有
	fmt.Println(pupil2.Student.Name) // haha
	fmt.Println(pupil2.Name)         // haha

	// 当结构体和匿名结构体有同名的字段或者方法采用就近访问的原则，如果希望访问匿名结构体的字段或者名字，
	// 通过匿名结构体名字区分
	pupil2.showHomeInfo()      // 访问的Pupil中的方法  pupil home info not support show
	pupil2.Home.showHomeInfo() // 访问的Home中的方法  home address=address
}

package main

import "fmt"

type Brand struct {
	Name    string
	Address string
}

type Good struct {
	Name  string
	Price float64
}

// 嵌套结构体指针

type Tv1 struct {
	*Good
	*Brand
}

func main() {
	tv1 := Tv1{
		Good: &Good{
			Name:  "tv001",
			Price: 1000.00,
		},
		Brand: &Brand{
			Name:    "haier",
			Address: "shandong",
		},
	}
	fmt.Println(tv1)                   // {0xc000008048 0xc00006e3c0}
	fmt.Println(*tv1.Good, *tv1.Brand) // {tv001 1000} {haier shandong}
	fmt.Println(tv1.Brand.Address)     // shandong
	fmt.Println(tv1.Price)             // 1000
}

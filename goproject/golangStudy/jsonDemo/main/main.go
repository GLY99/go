package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name string `json:"monster_name"` // : 左右不能有空格
	Age  int    `json:"monster_age"`
}

func main() {
	// 将结构体 map slice 序列化和反序列化
	monster := &Monster{Name: "sunwukong", Age: 5000}
	data, err := json.Marshal(monster)
	if err != nil {
		fmt.Printf("序列化失败, err=%v\n", err)
	} else {
		fmt.Printf("%v序列化后是%s\n", *monster, string(data)) // {sunwukong 5000}序列化后是{"monster_name":"sunwukong","monster_age":5000}
	}
	var newMonster Monster
	err = json.Unmarshal(data, &newMonster)
	if err != nil {
		fmt.Printf("反序列化失败, err=%v\n", err)
	} else {
		fmt.Printf("%v反序列化后是%v\n", data, newMonster)
		// [123 34 109 111 110 115 116 101 114 95 110 97 109 101 34 58 34 115 117 110 119 117 107 111
		//  110 103 34 44 34 109 111 110 115 116 101 114 95 97 103 101 34 58 53 48 48 48 125]反序
		// 列化后是{sunwukong 5000}
	}

	var myMap map[string]interface{} = map[string]interface{}{"name": "liming", "age": 18}
	data, err = json.Marshal(myMap)
	if err != nil {
		fmt.Printf("序列化失败, err=%v\n", err)
	} else {
		fmt.Printf("%v序列化后是%s\n", myMap, string(data)) // map[age:18 name:liming]序列化后是{"age":18,"name":"liming"}
	}
	var newMyMap map[string]interface{} // map在反序列化时不需要make,Unmarshal会自己make
	err = json.Unmarshal(data, &newMyMap)
	if err != nil {
		fmt.Printf("反序列化失败, err=%v\n", err)
	} else {
		fmt.Printf("%v反序列化后是%v\n", data, newMyMap)
		// [123 34 97 103 101 34 58 49 56 44 34 110 97 109 101 34 58 34 108 105 109 105 110 103 34 125]反序列化后是map[age:18 name:liming]
	}

	var slice []interface{} = []interface{}{"lili", 18}
	data, err = json.Marshal(slice)
	if err != nil {
		fmt.Printf("序列化失败, err=%v\n", err)
	} else {
		fmt.Printf("%v序列化后是%s\n", slice, string(data)) // [lili 18]序列化后是["lili",18]
	}
	var newSlice []interface{} // slice在反序列化时也不需要make,Unmarshal会自己make
	err = json.Unmarshal(data, &newSlice)
	if err != nil {
		fmt.Printf("反序列化失败, err=%v\n", err)
	} else {
		fmt.Printf("%v反序列化后是%v\n", data, newSlice) // [91 34 108 105 108 105 34 44 49 56 93]反序列化后是[lili 18]
	}

	// 对基本数据类型序列化,对基本数据类型序列化没有什么实际意义
	var num1 float64 = 1.234
	data, err = json.Marshal(num1)
	if err != nil {
		fmt.Printf("序列化失败, err=%v\n", err)
	} else {
		fmt.Printf("%v序列化后是%s\n", num1, string(data)) // 1.234序列化后是1.234
	}
}

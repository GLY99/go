package main

import (
	"fmt"
	"sort"
)

func main() {
	// 声明map不会分配内存，需要用make分配内存
	var map1 map[int]string
	map1 = make(map[int]string, 2)
	map1[1] = "a"
	map1[2] = "b"
	fmt.Println(map1)

	// map的使用
	// 先声明再make
	var map2 map[int]string
	map2 = make(map[int]string, 1)
	// 声明的时候直接make
	var map3 = make(map[int]string, 1)
	// 声明的时候直接赋值
	map4 := map[int]string{1: "A", 2: "B"}
	fmt.Println(map2, map3, map4)

	// value是map的map
	map5 := make(map[int]map[string]string, 5)
	map5[1] = make(map[string]string, 2)
	map5[1]["name"] = "tom"
	map5[1]["sex"] = "M"
	map5[2] = make(map[string]string, 2)
	map5[2]["name"] = "jerry"
	map5[2]["sex"] = "W"
	fmt.Println(map5)

	// map的更新和增加，一个key如果存在就是更新不存在就是增加
	map4[1] = "Z"
	fmt.Println(map4) // map[1:Z 2:B]

	// map的删除用的是一个内置函数
	delete(map4, 1)
	fmt.Println(map4) // map[2:B]

	// map的查找
	val, findRes := map4[2]
	fmt.Println(val, findRes) // B true
	val, findRes = map4[1]
	fmt.Println(val, findRes) // 空值 false

	// map遍历
	for k, v := range map4 {
		fmt.Println(k, v) // 2 B
	}
	for k1, v1 := range map5 {
		fmt.Println(k1, v1)
		// 1 map[name:tom sex:M]
		// 2 map[name:jerry sex:W]
		for k2, v2 := range v1 {
			fmt.Println(k2, v2)
			// name tom
			// sex M
			// name jerry
			// sex W
		}
	}

	// map长度
	fmt.Println(len(map5)) // 2

	// map切片
	monsters := make([]map[string]string, 2)
	fmt.Println(monsters) // [map[] map[]]
	monsters[0] = map[string]string{"name": "sunwukong"}
	monsters[1] = make(map[string]string, 1)
	monsters[1] = map[string]string{"name": "baigujing"}
	fmt.Println(monsters) // [map[name:sunwukong] map[name:baigujing]]
	// map切片大小固定，通过append进行动态扩容
	monsters = append(monsters, map[string]string{"name": "zhubajie"})
	fmt.Println(monsters) // [map[name:sunwukong] map[name:baigujing] map[name:zhubajie]]
	monster := map[string]string{"name": "tangseng"}
	monsters = append(monsters, monster)
	fmt.Println(monsters) // [map[name:sunwukong] map[name:baigujing] map[name:zhubajie] map[name:tangseng]]

	// map排序，先将key排序，然后按照key依次输出
	sortMap := make(map[int]string, 5)
	sortMap[100] = "Z"
	sortMap[1] = "X"
	sortMap[80] = "Y"
	idxSlice := make([]int, 0)
	for k, _ := range sortMap {
		idxSlice = append(idxSlice, k)
	}
	sort.Ints(idxSlice)
	for i := 0; i < len(idxSlice); i++ {
		fmt.Println(sortMap[idxSlice[i]])
	}
}
